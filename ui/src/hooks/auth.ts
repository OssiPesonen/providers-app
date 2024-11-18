'use client'
import { useState } from "react";
import { getLocalStorageItem, setLocalStorageItem  } from "@/util/localStorage";
import { jwtDecode } from "jwt-decode";
import { TrafficLightsServiceClient } from "@/proto/Traffic_lights_serviceServiceClientPb";
import {
  LoginRequest,
  RefreshTokenRequest,
} from "@/proto/traffic_lights_service_pb";

const accessTokenCacheKey = "access-token";
const refreshTokenCacheKey = "refresh-token";

const useAuth = (client: TrafficLightsServiceClient) => {
  // this is pretty primitive but it should suffice
  const [authError, setError] = useState<string | undefined>(undefined);

  const isTokenStillValid = (token: string) => {
    const payload = jwtDecode(token);
    const tokenExp = payload.exp ?? 0;
    const now = new Date().getTime() / 1000;
    const expired = tokenExp < now
    return !expired;
  };

  const refreshToken = () => {
    setError(undefined);

    const refreshToken = getLocalStorageItem(refreshTokenCacheKey);
    if (!refreshToken) {
      // Todo: Write a custom error
      throw new Error("Refresh token missing");
    }

    const req = new RefreshTokenRequest();
    req.setRefreshToken(refreshToken);

    client.refreshToken(req, {}, function (err, resp) {
      if (!err && typeof localStorage !== 'undefined') {
        setLocalStorageItem(refreshTokenCacheKey, resp.getRefreshToken());
        setLocalStorageItem(accessTokenCacheKey, resp.getAccessToken());
        setIsAuthenticated(true);
      } else {
        setError(err.message);
      }
    });
  };

  const getAccessToken = () => {
    const token: null | string = getLocalStorageItem(accessTokenCacheKey);
    // Token in cache but has expired
    if (token && !isTokenStillValid(token)) {
      // Attempt a refresh
      refreshToken();
      return getLocalStorageItem(accessTokenCacheKey);
    }
    return token;
  };

  const [isAuthenticated, setIsAuthenticated] = useState(
    getAccessToken() !== null
  );

  const login = (email: string, password: string) => {
    setError(undefined);
    const req = new LoginRequest();
    req.setEmail(email);
    req.setPassword(password);

    client.getToken(req, {}, function (err, resp) {
      if (!err) {
        localStorage.setItem(refreshTokenCacheKey, resp.getRefreshToken());
        localStorage.setItem(accessTokenCacheKey, resp.getAccessToken());
        setIsAuthenticated(true);
      } else {
        setError(err.message);
      }
    });
  };

  return {
    login,
    getAccessToken,
    isAuthenticated,
    authError,
  };
};

export { useAuth };
