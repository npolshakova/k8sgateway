package controller

import (
	"context"

	"github.com/solo-io/gloo/projects/gateway2/wellknown"

	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	"github.com/solo-io/gloo/projects/gateway2/controller/scheme"
	"github.com/solo-io/gloo/projects/gateway2/discovery"
	"github.com/solo-io/gloo/projects/gateway2/secrets"
	"github.com/solo-io/gloo/projects/gateway2/xds"
	"github.com/solo-io/gloo/projects/gloo/pkg/bootstrap"
	"github.com/solo-io/gloo/projects/gloo/pkg/syncer/sanitizer"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/healthz"
	metricsserver "sigs.k8s.io/controller-runtime/pkg/metrics/server"
	apiv1 "sigs.k8s.io/gateway-api/apis/v1"
)

const (
	// AutoProvision controls whether the controller will be responsible for provisioning dynamic
	// infrastructure for the Gateway API.
	AutoProvision = true
)

var (
	gatewayClass = apiv1.ObjectName(wellknown.GatewayClassName)

	setupLog = ctrl.Log.WithName("setup")
)

type StartConfig struct {
	Dev          bool
	ControlPlane bootstrap.ControlPlane
}

// Start runs the controllers responsible for processing the K8s Gateway API objects
// It is intended to be run in a goroutine as the function will block until the supplied
// context is cancelled
func Start(ctx context.Context, cfg StartConfig) error {
	var opts []zap.Opts
	if cfg.Dev {
		setupLog.Info("starting log in dev mode")
		opts = append(opts, zap.UseDevMode(true))
	}
	ctrl.SetLogger(zap.New(opts...))

	mgrOpts := ctrl.Options{
		Scheme:           scheme.NewScheme(),
		PprofBindAddress: "127.0.0.1:9099",
		// if you change the port here, also change the port "health" in the helmchart.
		HealthProbeBindAddress: ":9093",
		Metrics: metricsserver.Options{
			BindAddress: ":9092",
		},
	}
	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), mgrOpts)
	if err != nil {
		setupLog.Error(err, "unable to start manager")
		return err
	}

	// TODO: replace this with something that checks that we have xds snapshot ready (or that we don't need one).
	mgr.AddReadyzCheck("ready-ping", healthz.Ping)

	glooTranslator := newGlooTranslator(ctx)
	var sanz sanitizer.XdsSanitizers
	inputChannels := xds.NewXdsInputChannels()
	xdsSyncer := xds.NewXdsSyncer(
		wellknown.GatewayControllerName,
		glooTranslator,
		sanz,
		cfg.ControlPlane.SnapshotCache,
		false,
		inputChannels,
		mgr.GetClient(),
		mgr.GetScheme(),
	)
	if err := mgr.Add(xdsSyncer); err != nil {
		setupLog.Error(err, "unable to add xdsSyncer runnable")
		return err
	}

	gwCfg := GatewayConfig{
		Mgr:            mgr,
		GWClass:        gatewayClass,
		ControllerName: wellknown.GatewayControllerName,
		AutoProvision:  AutoProvision,
		ControlPlane:   cfg.ControlPlane,
		Kick:           inputChannels.Kick,
	}
	if err = NewBaseGatewayController(ctx, gwCfg); err != nil {
		setupLog.Error(err, "unable to create controller")
		return err
	}

	if err = discovery.NewDiscoveryController(ctx, mgr, inputChannels); err != nil {
		setupLog.Error(err, "unable to create controller")
		return err
	}

	if err = secrets.NewSecretsController(ctx, mgr, inputChannels); err != nil {
		setupLog.Error(err, "unable to create controller")
		return err
	}

	return mgr.Start(ctx)
}