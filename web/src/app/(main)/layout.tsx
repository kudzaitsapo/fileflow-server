import type { Metadata } from "next";
//import "@/styles/index.css";
import "@/app/globals.css";
import AuthProvider from "@/providers/auth";
import { AxiosProvider } from "@/providers/axios";
import Header from "@/components/header";
import SideBar from "@/components/sidebar";
import { ActiveProjectProvider } from "@/providers/project";
import React from "react";

export const metadata: Metadata = {
  title: "Fileflow Web: File Management at its best",
  description: "",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body>
        <AuthProvider>
          <AxiosProvider>
            <ActiveProjectProvider>
              <div className="bg-gray-100 text-gray-800">
                <Header />
                <div className="flex h-[calc(100vh-4rem)]">
                  <SideBar />
                  {children}
                </div>
              </div>
            </ActiveProjectProvider>
          </AxiosProvider>
        </AuthProvider>
      </body>
    </html>
  );
}
