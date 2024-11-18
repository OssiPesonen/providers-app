"use client";
import { useEffect, useState } from "react";
import google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb.js";
import styles from "./page.module.css";
import { TrafficLightsServiceClient } from "@/proto/Traffic_lights_serviceServiceClientPb";

import { useAuth } from "@/hooks/auth";

type Provider = {
  id: number;
  name: string;
  region: string;
  city: string;
  lob: string;
};

export default function Home() {
  const [providers, setProviders] = useState<Provider[]>([]);

  const client = new TrafficLightsServiceClient("http://localhost:8080");
  const { login, isAuthenticated, getAccessToken } = useAuth(client);

  const listProviders = () => {
    client.listProviders(
      new google_protobuf_empty_pb.Empty(),
      {
        Authorization: `Bearer ${getAccessToken()}`,
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
    if (!isAuthenticated) {
      // Todo: redirect to login instead of this
      login('demo@example.com', 'password1234');
    } else {
      // Todo: Move to some hook or repository
      listProviders();
    }
  }, [isAuthenticated]);

  return (
    <div className={styles.page}>
      <main className={styles.main}>
        <h3>Hello!</h3>
        <p>
          <strong>You are</strong>:{" "}
          {isAuthenticated ? "Authenticated" : "Unauthenticated" }
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
