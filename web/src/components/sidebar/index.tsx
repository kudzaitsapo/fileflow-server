"use client";
import React from "react";
import { signOut } from "next-auth/react";
import { usePathname, useRouter } from "next/navigation";
import {
  BarGraphIcon,
  ChatIcon,
  ExclamationCircleIcon,
  PlusIcon,
  SettingsIcon,
  ShieldIcon,
  SignOutIcon,
  UsersIcon,
} from "../icons";

const SideBar: React.FC = () => {
  const router = useRouter();
  const currentPageName = usePathname();

  const handleRedirect = (page: string) => {
    router.push(page);
  };

  const logOut = async () => {
    await signOut();
  };

  const activeClass = (page: string) =>
    currentPageName === page
      ? "bg-blue-500 text-white"
      : "text-gray-600 hover:bg-gray-100";

  const activeIcon = (page: string) =>
    currentPageName === page ? "text-white" : "text-gray-500";

  return (
    <>
      <div className="w-64 bg-white border-r border-gray-200 p-4 pt-6 flex flex-col gap-6">
        <button
          className="bg-blue-600 text-white border-none rounded-md py-3 px-4 font-semibold cursor-pointer flex items-center justify-center gap-2 transition-colors hover:bg-blue-800"
          onClick={() => handleRedirect("/projects/create")}
        >
          <PlusIcon />
          New Project
        </button>

        <div className="flex flex-col gap-2">
          <div className="text-xs font-semibold uppercase text-gray-500 px-2 mb-1">
            Project Management
          </div>
          <div className="overflow-y-auto flex-1">
            <div
              className={`flex items-center gap-3 p-2 rounded-md text-sm cursor-pointer ${activeClass(
                "/"
              )}`}
              onClick={() => {
                handleRedirect("/");
              }}
            >
              <svg
                className={`w-4 h-4 $${activeIcon("/")}`}
                xmlns="http://www.w3.org/2000/svg"
                viewBox="0 0 24 24"
                fill="none"
                stroke="currentColor"
                strokeWidth="2"
                strokeLinecap="round"
                strokeLinejoin="round"
              >
                <path d="M22 19a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h5l2 3h9a2 2 0 0 1 2 2z"></path>
              </svg>
              Files
            </div>
          </div>
        </div>

        <div className="flex flex-col gap-2">
          <div className="text-xs font-semibold uppercase text-gray-500 px-2 mb-1">
            Settings
          </div>
          <div
            className={`flex items-center gap-3 p-2 rounded-md text-sm cursor-pointer ${activeClass(
              "/users"
            )}`}
            onClick={() => {
              handleRedirect("/users");
            }}
          >
            <UsersIcon className={`w-4 h-4 ${activeIcon("/users")}`} />
            Users
          </div>
          <div className="flex items-center gap-3 p-2 rounded-md text-gray-600 text-sm cursor-pointer hover:bg-gray-100">
            <ChatIcon className="w-4 h-4 text-gray-500" />
            Activity Log
          </div>

          <div className="flex items-center gap-3 p-2 rounded-md text-gray-600 text-sm cursor-pointer hover:bg-gray-100">
            <BarGraphIcon className="w-4 h-4 text-gray-500" />
            Usage & Analytics
          </div>
          <div className="flex items-center gap-3 p-2 rounded-md text-gray-600 text-sm cursor-pointer hover:bg-gray-100">
            <SettingsIcon className="w-4 h-4 text-gray-500" />
            System Settings
          </div>
          <div className="flex items-center gap-3 p-2 rounded-md text-gray-600 text-sm cursor-pointer hover:bg-gray-100">
            <ShieldIcon className="w-4 h-4 text-gray-500" />
            Security
          </div>
          <div className="flex items-center gap-3 p-2 rounded-md text-gray-600 text-sm cursor-pointer hover:bg-gray-100">
            <ExclamationCircleIcon className="w-4 h-4 text-gray-500" />
            Help & Support
          </div>
        </div>

        <div className="flex flex-col gap-2">
          <div className="text-xs font-semibold uppercase text-gray-500 px-2 mb-1">
            Other
          </div>
          <div
            className="flex items-center gap-3 p-2 rounded-md text-gray-600 text-sm cursor-pointer hover:bg-gray-100"
            onClick={logOut}
          >
            <SignOutIcon className="w-4 h-4 text-gray-500" />
            Sign Out
          </div>
        </div>
      </div>
    </>
  );
};

export default SideBar;
