
---
title: "processing_mode.proto"
weight: 5
---

<!-- Code generated by solo-kit. DO NOT EDIT. -->


### Package: `solo.io.envoy.extensions.filters.http.ext_proc.v3`  
copied from https://github.com/envoyproxy/envoy/blob/ad89a587aa0177bfdad6b5c968a6aead5d9be7a4/api/envoy/extensions/filters/http/ext_proc/v3/processing_mode.proto


 
#### Types:


- [ProcessingMode](#processingmode)
- [HeaderSendMode](#headersendmode)
- [BodySendMode](#bodysendmode)
  



##### Source File: [github.com/solo-io/gloo/projects/controller/api/external/envoy/extensions/filters/http/ext_proc/v3/processing_mode.proto](https://github.com/solo-io/gloo/blob/main/projects/controller/api/external/envoy/extensions/filters/http/ext_proc/v3/processing_mode.proto)





---
### ProcessingMode

 
[#next-free-field: 7]

```yaml
"requestHeaderMode": .solo.io.envoy.extensions.filters.http.ext_proc.v3.ProcessingMode.HeaderSendMode
"responseHeaderMode": .solo.io.envoy.extensions.filters.http.ext_proc.v3.ProcessingMode.HeaderSendMode
"requestBodyMode": .solo.io.envoy.extensions.filters.http.ext_proc.v3.ProcessingMode.BodySendMode
"responseBodyMode": .solo.io.envoy.extensions.filters.http.ext_proc.v3.ProcessingMode.BodySendMode
"requestTrailerMode": .solo.io.envoy.extensions.filters.http.ext_proc.v3.ProcessingMode.HeaderSendMode
"responseTrailerMode": .solo.io.envoy.extensions.filters.http.ext_proc.v3.ProcessingMode.HeaderSendMode

```

| Field | Type | Description |
| ----- | ---- | ----------- | 
| `requestHeaderMode` | [.solo.io.envoy.extensions.filters.http.ext_proc.v3.ProcessingMode.HeaderSendMode](../processing_mode.proto.sk/#headersendmode) | How to handle the request header. Default is "SEND". |
| `responseHeaderMode` | [.solo.io.envoy.extensions.filters.http.ext_proc.v3.ProcessingMode.HeaderSendMode](../processing_mode.proto.sk/#headersendmode) | How to handle the response header. Default is "SEND". |
| `requestBodyMode` | [.solo.io.envoy.extensions.filters.http.ext_proc.v3.ProcessingMode.BodySendMode](../processing_mode.proto.sk/#bodysendmode) | How to handle the request body. Default is "NONE". |
| `responseBodyMode` | [.solo.io.envoy.extensions.filters.http.ext_proc.v3.ProcessingMode.BodySendMode](../processing_mode.proto.sk/#bodysendmode) | How do handle the response body. Default is "NONE". |
| `requestTrailerMode` | [.solo.io.envoy.extensions.filters.http.ext_proc.v3.ProcessingMode.HeaderSendMode](../processing_mode.proto.sk/#headersendmode) | How to handle the request trailers. Default is "SKIP". |
| `responseTrailerMode` | [.solo.io.envoy.extensions.filters.http.ext_proc.v3.ProcessingMode.HeaderSendMode](../processing_mode.proto.sk/#headersendmode) | How to handle the response trailers. Default is "SKIP". |




---
### HeaderSendMode

 
Control how headers and trailers are handled

| Name | Description |
| ----- | ----------- | 
| `DEFAULT` | The default HeaderSendMode depends on which part of the message is being processed. By default, request and response headers are sent, while trailers are skipped. |
| `SEND` | Send the header or trailer. |
| `SKIP` | Do not send the header or trailer. |




---
### BodySendMode

 
Control how the request and response bodies are handled

| Name | Description |
| ----- | ----------- | 
| `NONE` | Do not send the body at all. This is the default. |
| `STREAMED` | Stream the body to the server in pieces as they arrive at the proxy. |
| `BUFFERED` | Buffer the message body in memory and send the entire body at once. If the body exceeds the configured buffer limit, then the downstream system will receive an error. |
| `BUFFERED_PARTIAL` | Buffer the message body in memory and send the entire body in one chunk. If the body exceeds the configured buffer limit, then the body contents up to the buffer limit will be sent. |





<!-- Start of HubSpot Embed Code -->
<script type="text/javascript" id="hs-script-loader" async defer src="//js.hs-scripts.com/5130874.js"></script>
<!-- End of HubSpot Embed Code -->