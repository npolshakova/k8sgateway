// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/controller/api/v1/http_listener_options.proto

package v1

import (
	"encoding/binary"
	"errors"
	"fmt"
	"hash"
	"hash/fnv"
	"strconv"

	safe_hasher "github.com/solo-io/protoc-gen-ext/pkg/hasher"
	"github.com/solo-io/protoc-gen-ext/pkg/hasher/hashstructure"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = new(hash.Hash64)
	_ = fnv.New64
	_ = strconv.Itoa
	_ = hashstructure.Hash
	_ = new(safe_hasher.SafeHasher)
)

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *HttpListenerOptions) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("gloo.solo.io.github.com/solo-io/gloo/projects/controller/pkg/api/v1.HttpListenerOptions")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetGrpcWeb()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("GrpcWeb")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetGrpcWeb(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("GrpcWeb")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetHttpConnectionManagerSettings()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("HttpConnectionManagerSettings")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetHttpConnectionManagerSettings(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("HttpConnectionManagerSettings")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetHealthCheck()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("HealthCheck")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetHealthCheck(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("HealthCheck")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetExtensions()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Extensions")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetExtensions(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Extensions")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetWaf()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Waf")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetWaf(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Waf")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetDlp()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Dlp")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetDlp(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Dlp")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetWasm()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Wasm")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetWasm(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Wasm")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetExtauth()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Extauth")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetExtauth(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Extauth")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetRatelimitServer()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("RatelimitServer")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetRatelimitServer(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("RatelimitServer")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetCaching()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Caching")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetCaching(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Caching")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetGzip()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Gzip")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetGzip(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Gzip")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetProxyLatency()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("ProxyLatency")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetProxyLatency(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("ProxyLatency")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetBuffer()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Buffer")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetBuffer(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Buffer")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetCsrf()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Csrf")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetCsrf(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Csrf")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetGrpcJsonTranscoder()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("GrpcJsonTranscoder")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetGrpcJsonTranscoder(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("GrpcJsonTranscoder")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetSanitizeClusterHeader()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("SanitizeClusterHeader")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetSanitizeClusterHeader(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("SanitizeClusterHeader")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetLeftmostXffAddress()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("LeftmostXffAddress")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetLeftmostXffAddress(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("LeftmostXffAddress")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetDynamicForwardProxy()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("DynamicForwardProxy")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetDynamicForwardProxy(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("DynamicForwardProxy")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetConnectionLimit()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("ConnectionLimit")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetConnectionLimit(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("ConnectionLimit")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetNetworkLocalRatelimit()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("NetworkLocalRatelimit")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetNetworkLocalRatelimit(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("NetworkLocalRatelimit")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetHttpLocalRatelimit()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("HttpLocalRatelimit")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetHttpLocalRatelimit(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("HttpLocalRatelimit")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetRouter()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Router")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetRouter(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Router")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetTap()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Tap")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetTap(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Tap")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetStatefulSession()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("StatefulSession")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetStatefulSession(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("StatefulSession")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetHeaderValidationSettings()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("HeaderValidationSettings")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetHeaderValidationSettings(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("HeaderValidationSettings")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	switch m.ExtProcConfig.(type) {

	case *HttpListenerOptions_DisableExtProc:

		if h, ok := interface{}(m.GetDisableExtProc()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("DisableExtProc")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetDisableExtProc(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("DisableExtProc")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *HttpListenerOptions_ExtProc:

		if h, ok := interface{}(m.GetExtProc()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("ExtProc")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetExtProc(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("ExtProc")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	return hasher.Sum64(), nil
}