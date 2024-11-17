"use client";
import styles from "./page.module.css";
import { useEffect, useState } from "react";
import { TrafficLightsServiceClient } from "@/proto/Traffic_lights_serviceServiceClientPb";
import google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb.js";
import { LoginRequest } from "@/proto/traffic_lights_service_pb";

type Provider = {
  id: number;
  name: string;
  region: string;
  city: string;
  lob: string;
};

export default function Home() {
  const [isUnauthenticated, setIsUnauthenticated] = useState<
    boolean | undefined
  >(undefined);
  const [providers, setProviders] = useState<Provider[]>([]);

  const client = new TrafficLightsServiceClient("http://localhost:8080");

  const login = () => {
    if (localStorage.getItem("access-token") !== null) {
      return;
    }

    const req = new LoginRequest();
    req.setEmail("demo@example.com");
    req.setPassword("password1234");

    client.getToken(req, {}, function (err, resp) {
      setIsUnauthenticated(!err);

      if (!err) {
        localStorage.setItem("refresh-token", resp.getRefreshToken());
        localStorage.setItem("access-token", resp.getAccessToken());
      }
    });
  };

  const listProviders = () => {
    client.listProviders(
      new google_protobuf_empty_pb.Empty(),
      {
        Authorization: `Bearer ${localStorage.getItem("access-token")}`,
      },
      function (err, resp) {
        if (err) {
          console.error(err);
        } else {
          const providersList = resp.getProvidersList();
          const providers: Provider[] = [];
          providersList.forEach((p) => {
            providers.push({
              id: p.getId(),
              name: p.getName(),
              city: p.getCity(),
              region: p.getRegion(),
              lob: p.getLineOfBusiness(),
            });
          });
          setProviders(providers);
        }
      }
    );
  };

  useEffect(() => {
    if (localStorage.getItem("access-token") !== null) {
      setIsUnauthenticated(true);
    } else {
      login();
    }
  }, []);

  useEffect(() => {
    if (isUnauthenticated) {
      listProviders();
    }
  }, [isUnauthenticated]);

  return (
    <div className={styles.page}>
      <main className={styles.main}>
        <h3>Hello!</h3>
        <p>
          <strong>You are</strong>:{" "}
          {isUnauthenticated === undefined
            ? "Wait a second.."
            : isUnauthenticated === false
            ? "Unauthenticated :("
            : "Authenticated!"}
        </p>
        <p>Here is a list of providers in the system:</p>
        <ol>
          {providers.map((p) => (
            <li key={p.id}>
              {p.name}, {p.region}, {p.city}
            </li>
          ))}
        </ol>
      </main>
    </div>
  );
}
