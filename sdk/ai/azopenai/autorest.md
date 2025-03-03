# Go

These settings apply only when `--go` is specified on the command line.

``` yaml
input-file:
#- https://raw.githubusercontent.com/Azure/azure-rest-api-specs/13a645b66b741e3cc2ef378cb81974b30e6a7a86/specification/cognitiveservices/AzureOpenAI/inference/2023-06-01-preview/generated.json
- ./testdata/generated/openapi3.json

output-folder: ../azopenai
clear-output-folder: false
module: github.com/Azure/azure-sdk-for-go/sdk/ai/azopenai
license-header: MICROSOFT_MIT_NO_VERSION
openapi-type: data-plane
go: true
use: "@autorest/go@4.0.0-preview.52"
title: "OpenAI"
slice-elements-byval: true
# can't use this since it removes an innererror type that we want ()
# remove-non-reference-schema: true
```

## Transformations

``` yaml
directive:
  # Add x-ms-parameter-location to parameters in x-ms-parameterized-host
  - from: openapi-document
    where: $.servers.0.variables.endpoint
    debug: true
    transform: $["x-ms-parameter-location"] = "client";

  # Make deploymentId a client parameter
  # This must be done in each operation as the parameter is not defined in the components section
  - from: openapi-document
    where: $.paths..parameters..[?(@.name=='deploymentId')]
    transform: $["x-ms-parameter-location"] = "client";

  - from: openapi-document
    where: $..paths["/deployments/{deploymentId}/completions"].post.requestBody
    transform: $["required"] = true;
  - from: openapi-document
    where: $.paths["/deployments/{deploymentId}/embeddings"].post.requestBody
    transform: $["required"] = true;

  # get rid of these auto-generated LRO status methods that aren't exposed.
  - from: openapi-document
    where: $.paths
    transform: delete $["/operations/images/{operationId}"]

  # Remove stream property from CompletionsOptions and ChatCompletionsOptions
  - from: openapi-document
    where: $.components.schemas["CompletionsOptions"]
    transform: delete $.properties.stream;
  - from: openapi-document
    where: $.components.schemas["ChatCompletionsOptions"]
    transform: delete $.properties.stream; 

  # Replace anyOf schemas with an empty schema (no type) to get an "any" type generated
  - from: openapi-document
    where: '$.components.schemas["EmbeddingsOptions"].properties["input"]'
    transform: delete $.anyOf;

  - from: openapi-document
    where: $.paths["/images/generations:submit"].post
    transform: $["x-ms-long-running-operation"] = true;

  # Fix autorest bug
  - from: openapi-document
    where: $.components.schemas["BatchImageGenerationOperationResponse"].properties
    transform: |
      $.result["$ref"] = "#/components/schemas/ImageGenerations"; delete $.allOf;
      $.status["$ref"] = "#/components/schemas/AzureOpenAIOperationState"; delete $.allOf;
      $.error["$ref"] = "#/components/schemas/Azure.Core.Foundations.Error"; delete $.allOf;
  - from: openapi-document
    where: $.components.schemas["ChatMessage"].properties.role
    transform: $["$ref"] = "#/components/schemas/ChatRole"; delete $.oneOf;
  - from: openapi-document
    where: $.components.schemas["Choice"].properties.finish_reason
    transform: $["$ref"] = "#/components/schemas/CompletionsFinishReason"; delete $.oneOf;
  - from: openapi-document
    where: $.components.schemas["ImageOperation"].properties.status
    transform: $["$ref"] = $.anyOf[0]["$ref"];delete $.anyOf;
  - from: openapi-document
    where: $.components.schemas.ImageGenerationOptions.properties
    transform: |
      $.size["$ref"] = "#/components/schemas/ImageSize"; delete $.allOf;
      $.response_format["$ref"] = "#/components/schemas/ImageGenerationResponseFormat"; delete $.allOf;
  - from: openapi-document
    where: $.components.schemas["ImageOperationResponse"].properties
    transform: |
      $.status["$ref"] = "#/components/schemas/State"; delete $.status.allOf;
      $.result["$ref"] = "#/components/schemas/ImageResponse"; delete $.status.allOf;      
  - from: openapi-document
    where: $.components.schemas["ImageOperationStatus"].properties.status
    transform: $["$ref"] = "#/components/schemas/State"; delete $.allOf; 
  - from: openapi-document
    where: $.components.schemas["ContentFilterResult"].properties.severity
    transform: $["$ref"] = "#/components/schemas/ContentFilterSeverity"; delete $.allOf;
  - from: openapi-document
    where: $.components.schemas["ChatChoice"].properties.finish_reason
    transform: $["$ref"] = "#/components/schemas/CompletionsFinishReason"; delete $.oneOf;
  - from: openapi-document
    where: $.components.schemas["AzureChatExtensionConfiguration"].properties.type
    transform: $["$ref"] = "#/components/schemas/AzureChatExtensionType"; delete $.allOf;
  - from: openapi-document
    where: $.components.schemas["AzureChatExtensionConfiguration"].properties.type
    transform: $["$ref"] = "#/components/schemas/AzureChatExtensionType"; delete $.allOf;
  - from: openapi-document
    where: $.components.schemas["AzureCognitiveSearchChatExtensionConfiguration"].properties.queryType
    transform: $["$ref"] = "#/components/schemas/AzureCognitiveSearchQueryType"; delete $.allOf;
  # Fix "AutoGenerated" models
  - from: openapi-document
    where: $.components.schemas["ChatCompletions"].properties.usage
    transform: >
      delete $.allOf;
      $["$ref"] = "#/components/schemas/CompletionsUsage";
  - from: openapi-document
    where: $.components.schemas["Completions"].properties.usage
    transform: >
      delete $.allOf;
      $["$ref"] = "#/components/schemas/CompletionsUsage";

  #
  # strip out the deploymentID validation code - we absorbed this into the endpoint.
  #
	# urlPath := "/deployments/{deploymentId}/embeddings"
	# if client.deploymentID == "" {
	# 	return nil, errors.New("parameter client.deploymentID cannot be empty")
	# }
	# urlPath = strings.ReplaceAll(urlPath, "{deploymentId}", url.PathEscape(client.deploymentID))
  - from: client.go
    where: $
    transform: >-
      return $.replace(
        /(\s+)urlPath\s*:=\s*"\/deployments\/\{deploymentId\}\/([^"]+)".+?url\.PathEscape.+?\n/gs, 
        "$1urlPath := \"$2\"\n")

  # Unexport the the poller state enum.
  - from: 
      - constants.go
      - models.go
    where: $
    transform: return $.replace(/AzureOpenAIOperationState/g, "azureOpenAIOperationState");

  # splice out the auto-generated `deploymentID` field from the client
  - from: client.go
    where: $
    transform: >-
      return $.replace(
        /(type Client struct[^}]+})/s, 
        "type Client struct {\ninternal *azcore.Client; clientData;\n}")

  - from: 
    - models_serde.go
    - models.go
    where: $
    transform: return $.replace(/AzureCoreFoundations/g, "azureCoreFoundations");
  - from: 
    - models_serde.go
    - models.go
    where: $
    transform: return $.replace(/(?:\/\/.*\s)?func \(\w \*?(?:ErrorResponse|ErrorResponseError|InnerError|InnerErrorInnererror)\).*\{\s(?:.+\s)+\}\s/g, "");

  - from: constants.go
    where: $
    transform: >-
      return $.replace(
        /type ServiceAPIVersions string.+PossibleServiceAPIVersionsValues.+?\n}/gs, 
        "")

  # delete client name prefix from method options and response types
  - from:
      - client.go
      - models.go
      - options.go
      - response_types.go
    where: $
    transform: return $.replace(/Client(\w+)((?:Options|Response))/g, "$1$2");

  # allow interception of formatting the URL path
  - from: client.go
    where: $
    transform: |
      return $
        .replace(/runtime\.JoinPaths\(client.endpoint, urlPath\)/g, "client.formatURL(urlPath, getDeployment(body))");

  # Some ImageGenerations hackery to represent the ImageLocation/ImagePayload polymorphism.
  # - Remove the auto-generated ImageGenerationsDataItem.
  # - Replace the ImageGenerations.Data type with []ImageGenerationDataItem
  # - from: models.go
  #   where: $
  #   transform: |
  #     return $.replace(/type ImageGenerationsDataItem struct {[^}]+}/, "// ImageGenerationsDataItem represents an image URL or payload\ntype ImageGenerationsDataItem struct{\nImageLocation\nImagePayload\n}")
  #       $.replace(/(type ImageGenerations struct.+?)Data any/g, "$1Data []ImageGenerationsDataItem")

  - from: models.go
    where: $
    transform: |
      return $.replace(/(type ImageGenerations struct.+?)Data any/sg, "$1Data []ImageGenerationsDataItem")

  # delete the auto-generated ImageGenerationsDataItem, we handle that custom
  - from: models.go
    where: $
    transform: return $.replace(/\/\/ ImageGenerationsDataItem represents[^}]+}/s, "");

  # rename the image constants
  - from: constants.go
    where: $
    transform: |
      return $.replace(/ImageSizeFiveHundredTwelveX512/g, "ImageSize512x512")
        .replace(/ImageSizeOneThousandTwentyFourX1024/g, "ImageSize1024x1024")
        .replace(/ImageSizeTwoHundredFiftySixX256/g, "ImageSize256x256");

  # scrub the Image(Payload|Location) deserializers.
  - from: models_serde.go
    where: $
    transform: |
      return $.replace(/\/\/ UnmarshalJSON implements the json.Unmarshaller interface for type ImagePayload.+?\n}/s, "")
        .replace(/\/\/ MarshalJSON implements the json.Marshaller interface for type ImagePayload.+?\n}/s, "")
        .replace(/\/\/ UnmarshalJSON implements the json.Unmarshaller interface for type ImageLocation.+?\n}/s, "")
        .replace(/\/\/ MarshalJSON implements the json.Marshaller interface for type ImageLocation.+?\n}/s, "");

  # hide the image generation pollers.
  - rename-operation:
      from: beginAzureBatchImageGeneration
      to: azureBatchImageGenerationInternal
  - from: 
    - client.go
    - models.go
    - models_serde.go
    - options.go
    - response_types.go
    where: $
    transform: |
      return $.replace(/GetAzureBatchImageGenerationOperationStatusResponse/g, "getAzureBatchImageGenerationOperationStatusResponse")
        .replace(/AzureBatchImageGenerationInternalResponse/g, "azureBatchImageGenerationInternalResponse")
        .replace(/GetAzureBatchImageGenerationOperationStatusOptions/g, "getAzureBatchImageGenerationOperationStatusOptions")
        .replace(/GetAzureBatchImageGenerationOperationStatus/g, "getAzureBatchImageGenerationOperationStatus")
        .replace(/BeginAzureBatchImageGenerationInternal/g, "beginAzureBatchImageGeneration")
        .replace(/BatchImageGenerationOperationResponse/g, "batchImageGenerationOperationResponse");

  # BUG: ChatCompletionsOptionsFunctionCall is another one of those "here's mutually exclusive values" options...
  - from: 
    - models.go
    - models_serde.go
    where: $
    transform: |
      return $
        .replace(/populateAny\(objectMap, "function_call", c.FunctionCall\)/, 'populate(objectMap, "function_call", c.FunctionCall)')
        .replace(/\/\/ ChatCompletionsOptionsFunctionCall.+?\n}/, "")
        .replace(/FunctionCall any/, "FunctionCall *ChatCompletionsOptionsFunctionCall");
  
  # fix some casing
  - from: 
    - client.go
    - models.go
    - models_serde.go
    - options.go
    - response_types.go
    where: $
    transform: return $.replace(/Logprobs/g, "LogProbs")

  # delete ContentFilterResult in favor of our custom representation.
  - from: 
    - models.go
    - models_serde.go
    where: $
    transform: |
      return $.replace(/\/\/ ContentFilterResult.+?\n}/s, "")
        .replace(/\/\/ MarshalJSON implements the json.Marshaller interface for type ContentFilterResult.+?\n}/s, "")
        .replace(/\/\/ UnmarshalJSON implements the json.Unmarshaller interface for type ContentFilterResult.+?\n}/s, "");

  - from: constants.go
    where: $
    transform: return $.replace(/\/\/ PossibleazureOpenAIOperationStateValues returns.+?\n}/s, "");

  # fix incorrect property name for content filtering
  # TODO: I imagine we should able to fix this in the tsp?
  - from: models_serde.go
    where: $
    transform: |
      return $
        .replace(/		case "selfHarm":/g, '		case "self_harm":')
        .replace(/populate\(objectMap, "selfHarm", c.SelfHarm\)/g, 'populate(objectMap, "self_harm", c.SelfHarm)');

  - from: client.go
    where: $
    transform: return $.replace(/runtime\.NewResponseError/sg, "client.newError");

  #
  # rename `Model` to `Deployment`
  #
  - from: models.go
    where: $
    transform: |
      return $
        .replace(/\/\/ The model name.*?Model \*string/sg, "// REQUIRED: Deployment specifies the name of the deployment (for Azure OpenAI) or model (for OpenAI) to use for this request.\nDeployment string");

  - from: models_serde.go
    where: $
    transform: |
      return $
        .replace(/populate\(objectMap, "model", (c|e).Model\)/g, 'populate(objectMap, "model", &$1.Deployment)')
        .replace(/err = unpopulate\(val, "Model", &(c|e).Model\)/g, 'err = unpopulate(val, "Model", &$1.Deployment)');

  # Make the Azure extensions internal - we expose these through the GetChatCompletions*() functions
  # and just treat which endpoint we use as an implementation detail.
  - from: client.go
    where: $
    transform: |
      return $
        .replace(/GetChatCompletionsWithAzureExtensions([ (])/g, "getChatCompletionsWithAzureExtensions$1")
        .replace(/GetChatCompletions([ (])/g, "getChatCompletions$1");

  # move the Azure extensions options into place
  - from: models.go
    where: $
    transform: return $.replace(/(\/\/ The configuration entries for Azure OpenAI.+?)DataSources \[\]AzureChatExtensionConfiguration/s, "$1AzureExtensionsOptions *AzureChatExtensionOptions");
  - from: models_serde.go
    where: $
    transform: |
      return $
        .replace(/populate\(objectMap, "dataSources", c.DataSources\)/, 'if c.AzureExtensionsOptions != nil { populate(objectMap, "dataSources", c.AzureExtensionsOptions.Extensions) }')
        // technically not used, but let's be completionists...
        .replace(/err = unpopulate\(val, "DataSources", &c.DataSources\)/, 'c.AzureExtensionsOptions = &AzureChatExtensionOptions{}; err = unpopulate(val, "DataSources", &c.AzureExtensionsOptions.Extensions)')

  # try to fix some of the generated types.

  # swap the `Parameters` and `Type` fields (Type really drives what's in Parameters)
  - from: models.go
    where: $
    transform: |
      let typeRE = /(\/\/ REQUIRED; The label for the type of an Azure chat extension.*?Type \*AzureChatExtensionType)/s;
      let paramsRE = /(\/\/ REQUIRED; The configuration payload used for the Azure chat extension.*?Parameters any)/s;

      return $
        .replace(paramsRE, "")
        .replace(typeRE, $.match(typeRE)[1] + "\n\n" + $.match(paramsRE)[1]);

  - from: constants.go
    where: $
    transform: |
      return $.replace(
        /(AzureChatExtensionTypeAzureCognitiveSearch AzureChatExtensionType)/, 
        "// AzureChatExtensionTypeAzureCognitiveSearch enables the use of an Azure Cognitive Search index with chat completions.\n// [AzureChatExtensionConfiguration.Parameter] should be of type [AzureCognitiveSearchChatExtensionConfiguration].\n$1");
```
