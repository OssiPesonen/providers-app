import { ProvidersAppServiceClient } from "$lib/proto/providers_app_service.client";
import { GrpcWebFetchTransport } from "@protobuf-ts/grpcweb-transport";

export const apiClient = () => {
    const transport = new GrpcWebFetchTransport({
        baseUrl: "http://localhost:8080"
    });

    return new ProvidersAppServiceClient(transport);
}