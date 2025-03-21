"use client";
import { Project } from "@/models/project";
import { useAxios } from "@/providers/axios";
import React, { useEffect } from "react";
import SidebarItem from "../sidebar-item";
import { signOut, useSession } from "next-auth/react";
import { useActiveProject } from "@/providers/project";
import { useRouter } from "next/navigation";
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
  const { get } = useAxios();
  const { data: session } = useSession();
  const [projects, setProjects] = React.useState<Project[]>([]);
  const { activeProject, setActiveProject } = useActiveProject();
  const router = useRouter();
  //const queryParams = useSearchParams();

  useEffect(() => {
    const fetchProjects = async () => {
      if (session && session.user) {
        const response = await get<Project[]>("/projects");
        return response;
      }
    };

    fetchProjects().then((apiProjects) => {
      if (Array.isArray(apiProjects)) {
        setProjects(apiProjects);
        setActiveProject(apiProjects[0]);
      } else {
        console.log("LOG::error fetching projects: ", apiProjects);
      }
    });
  }, [get, session, setActiveProject]);

  const handleRedirect = (page: string) => {
    router.push(page);
  };

  const handleProjectClick = (index: number) => {
    const project = projects.find((p) => p.id === index);
    if (!project) return;
    setActiveProject(project);
  };

  const logOut = async () => {
    await signOut();
  };

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
            Your Projects
          </div>
          <div className="overflow-y-auto flex-1">
            {projects &&
              projects.map((project) => (
                <SidebarItem
                  name={project.name}
                  onClick={() => handleProjectClick(project.id)}
                  isActive={activeProject?.id == project.id}
                  key={project.id}
                />
              ))}
          </div>
        </div>

        <div className="flex flex-col gap-2">
          <div className="text-xs font-semibold uppercase text-gray-500 px-2 mb-1">
            Settings
          </div>
          <div className="flex items-center gap-3 p-2 rounded-md text-gray-600 text-sm cursor-pointer hover:bg-gray-100">
            <UsersIcon className="w-4 h-4 text-gray-500" />
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
            Account Settings
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
