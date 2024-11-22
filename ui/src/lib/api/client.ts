import { TrafficLightsServiceClient } from "$lib/proto/traffic_lights_service.client";
import { GrpcWebFetchTransport } from "@protobuf-ts/grpcweb-transport";

export const apiClient = () => {
    const transport = new GrpcWebFetchTransport({
        baseUrl: "http://localhost:8080"
    });

    return new TrafficLightsServiceClient(transport);
}