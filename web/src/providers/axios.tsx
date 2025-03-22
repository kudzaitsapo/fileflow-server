"use client";

import React, { createContext, useContext, ReactNode } from "react";
import axios, { AxiosInstance } from "axios";
import { useSession } from "next-auth/react";
import { ApiError } from "@/models/error";

interface AxiosContextType {
  axiosInstance: AxiosInstance;

  get: <T>(
    url: string,
    pathParams?: Record<string, string>,
    headers?: Record<string, string>
  ) => Promise<T | ApiError>;

  post: <T>(
    url: string,
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    data: any,
    headers?: Record<string, string>,
    pathParams?: Record<string, string>
  ) => Promise<T>;

  put: <T>(
    url: string,
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    data: any,
    headers?: Record<string, string>,
    pathParams?: Record<string, string>
  ) => Promise<T>;

  delete: (
    url: string,
    pathParams?: Record<string, string>,
    headers?: Record<string, string>
  ) => Promise<void>;
}

const AxiosContext = createContext<AxiosContextType | null>(null);

interface AxiosProviderProps {
  children: ReactNode;
}

export const AxiosProvider = ({ children }: AxiosProviderProps) => {
  const { data: session } = useSession();

  // Create an axios instance (you may add default configurations here)
  const axiosInstance = axios.create({
    baseURL: process.env.NEXT_PUBLIC_BACKEND_URL || "", // update as needed
  });

  // GET method: retrieves data from the given URL
  const get = async <T,>(
    url: string,
    pathParams?: Record<string, string>,
    headers?: Record<string, string>
  ): Promise<T | ApiError> => {
    if (!session || !session.user || !session.user.accessToken) {
      throw new Error("No session available. Please login.");
    }
    try {
      const response = await axiosInstance.get<T>(url, {
        params: pathParams,
        headers: {
          Authorization: `Bearer ${session.user.accessToken}`,
          ...headers,
        },
      });
      return response.data;
    } catch (e) {
      console.log("LOG::axios-error: ", e);
      const err: ApiError = {
        message: "There was an error: ",
      };
      return err;
    }
  };

  const post = async <T,>(
    url: string,
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    data: any,
    headers?: Record<string, string>,
    pathParams?: Record<string, string>
  ): Promise<T> => {
    if (!session || !session.user || !session.user.accessToken) {
      throw new Error("No session available. Please login.");
    }
    const response = await axiosInstance.post<T>(url, data, {
      params: pathParams,
      headers: {
        ...headers,
        Authorization: `Bearer ${session.user.accessToken}`,
      },
    });
    return response.data;
  };

  const put = async <T,>(
    url: string,
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    data: any,
    headers?: Record<string, string>,
    pathParams?: Record<string, string>
  ): Promise<T> => {
    if (!session || !session.user || !session.user.accessToken) {
      throw new Error("No session available. Please login.");
    }
    const response = await axiosInstance.put<T>(url, data, {
      params: pathParams,
      headers: {
        ...headers,
        Authorization: `Bearer ${session.user.accessToken}`,
      },
    });
    return response.data;
  };

  const deleteMethod = async (
    url: string,
    pathParams?: Record<string, string>,
    headers?: Record<string, string>
  ): Promise<void> => {
    if (!session || !session.user || !session.user.accessToken) {
      throw new Error("No session available. Please login.");
    }
    await axiosInstance.delete(url, {
      params: pathParams,
      headers: {
        Authorization: `Bearer ${session.user.accessToken}`,
        ...headers,
      },
    });
  };

  return (
    <AxiosContext.Provider
      value={{ axiosInstance, get, post, put, delete: deleteMethod }}
    >
      {children}
    </AxiosContext.Provider>
  );
};

export const useAxios = () => {
  const context = useContext(AxiosContext);
  if (!context) {
    throw new Error("useAxiosContext must be used within an AxiosProvider");
  }
  return context;
};
