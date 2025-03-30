"use client";
import React, { useEffect } from "react";
import { signOut, useSession } from "next-auth/react";
import { useAxios } from "@/providers/axios";
import { Project } from "@/models/project";
import { useActiveProject } from "@/providers/project";
import { SignOutIcon } from "../icons";

const Header: React.FC = () => {
  const { data: session } = useSession();
  const { get } = useAxios();
  const [projects, setProjects] = React.useState<Project[]>([]);
  const [isDropdownOpen, setIsDropdownOpen] = React.useState<boolean>(false);
  const { activeProject, setActiveProject } = useActiveProject();

  useEffect(() => {
    const fetchProjects = async () => {
      const response = await get<Project[]>("/get-user-projects");
      if (!response.success) {
        throw new Error("Failed to fetch projects");
      }

      setProjects(response.result);
      if (!activeProject) {
        const firstProject = response.result[0];
        setActiveProject(firstProject);
      }
    };
    if (session && session.user) {
      fetchProjects();
    }
  }, [activeProject, get, session, setActiveProject]);

  const logOut = async () => {
    await signOut();
  };

  return (
    <>
      <header className="bg-white shadow sticky top-0 z-10 flex items-center justify-between px-6 py-4">
        {/* Left section with logo and project selection */}
        <div className="flex items-center gap-6">
          {/* Logo */}
          <div className="flex items-center gap-2 font-bold text-xl text-blue-800">
            <span className="w-8 h-8 bg-blue-600 rounded flex items-center justify-center text-white font-bold">
              F
            </span>
            <span>FileFlow</span>
          </div>

          {/* Project Selector with label */}
          <div className="flex items-center gap-2">
            <label
              htmlFor="project-selector"
              className="text-sm font-medium text-gray-700"
            >
              Project:
            </label>
            <div className="relative">
              <select
                id="project-selector"
                className="appearance-none bg-white border border-gray-300 text-gray-700 py-2 px-4 pr-8 rounded-md hover:border-blue-500 focus:outline-none focus:ring-1 focus:ring-blue-500 focus:border-blue-500 cursor-pointer font-medium min-w-[180px]"
                onChange={(e: React.ChangeEvent<HTMLSelectElement>) => {
                  const selectedProject = projects.find(
                    (project) => project.id === parseInt(e.target.value)
                  );
                  if (selectedProject) {
                    setActiveProject(selectedProject);
                  }
                }}
                value={activeProject ? activeProject.id : ""}
              >
                {projects.map((project) => (
                  <option key={project.id} value={project.id}>
                    {project.name}
                  </option>
                ))}
              </select>
              <div className="pointer-events-none absolute inset-y-0 right-0 flex items-center px-2 text-gray-700">
                <svg
                  className="fill-current h-4 w-4"
                  xmlns="http://www.w3.org/2000/svg"
                  viewBox="0 0 20 20"
                >
                  <path d="M9.293 12.95l.707.707L15.657 8l-1.414-1.414L10 10.828 5.757 6.586 4.343 8z" />
                </svg>
              </div>
            </div>
          </div>
        </div>

        {/* Right section with notifications, user info and dropdown menu */}
        <div className="flex items-center gap-5">
          {/* Notifications */}
          <button className="text-gray-600 hover:text-blue-600 focus:outline-none">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              className="h-6 w-6"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
            >
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                strokeWidth={2}
                d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9"
              />
            </svg>
          </button>

          {/* User profile section */}
          <div className="flex items-center gap-3">
            <span className="font-medium text-gray-800">
              {(session && session.user && session.user.name) ??
                "FileFlow User"}
            </span>
            <div className="w-9 h-9 rounded-full bg-blue-100 flex items-center justify-center text-blue-700 font-semibold shadow-sm border border-blue-200">
              {(session && session.user && session.user.name[0]) ?? `FF`}
            </div>

            {/* Dropdown menu */}
            <div className="relative">
              <button
                className="text-gray-600 hover:text-blue-600 focus:outline-none mt-3"
                onClick={() => setIsDropdownOpen(!isDropdownOpen)}
              >
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  className="h-5 w-5"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth={2}
                    d="M19 9l-7 7-7-7"
                  />
                </svg>
              </button>

              {/* Dropdown content */}
              {isDropdownOpen && (
                <div className="absolute right-0 mt-2 w-48 bg-white rounded-md shadow-lg py-1 z-20 border border-gray-200">
                  <a
                    href="#"
                    className="px-4 py-2 text-sm text-gray-700 hover:bg-blue-50 flex items-center"
                    onClick={() => setIsDropdownOpen(false)}
                  >
                    <svg
                      xmlns="http://www.w3.org/2000/svg"
                      className="h-4 w-4 mr-2 text-gray-500"
                      viewBox="0 0 20 20"
                      fill="currentColor"
                    >
                      <path
                        fillRule="evenodd"
                        d="M10 9a3 3 0 100-6 3 3 0 000 6zm-7 9a7 7 0 1114 0H3z"
                        clipRule="evenodd"
                      />
                    </svg>
                    My Profile
                  </a>
                  <a
                    href="#"
                    className="px-4 py-2 text-sm text-gray-700 hover:bg-blue-50 flex items-center"
                    onClick={() => setIsDropdownOpen(false)}
                  >
                    <svg
                      xmlns="http://www.w3.org/2000/svg"
                      className="h-4 w-4 mr-2 text-gray-500"
                      viewBox="0 0 20 20"
                      fill="currentColor"
                    >
                      <path
                        fillRule="evenodd"
                        d="M11.49 3.17c-.38-1.56-2.6-1.56-2.98 0a1.532 1.532 0 01-2.286.948c-1.372-.836-2.942.734-2.106 2.106.54.886.061 2.042-.947 2.287-1.561.379-1.561 2.6 0 2.978a1.532 1.532 0 01.947 2.287c-.836 1.372.734 2.942 2.106 2.106a1.532 1.532 0 012.287.947c.379 1.561 2.6 1.561 2.978 0a1.533 1.533 0 012.287-.947c1.372.836 2.942-.734 2.106-2.106a1.533 1.533 0 01.947-2.287c1.561-.379 1.561-2.6 0-2.978a1.532 1.532 0 01-.947-2.287c.836-1.372-.734-2.942-2.106-2.106a1.532 1.532 0 01-2.287-.947zM10 13a3 3 0 100-6 3 3 0 000 6z"
                        clipRule="evenodd"
                      />
                    </svg>
                    System Settings
                  </a>
                  <div className="border-t border-gray-100 my-1"></div>
                  <a
                    href="#"
                    className="px-4 py-2 text-sm text-red-600 hover:bg-red-50 flex items-center"
                    onClick={logOut}
                  >
                    <SignOutIcon className="h-4 w-4 mr-2 text-red-500" />
                    Sign Out
                  </a>
                </div>
              )}
            </div>
          </div>
        </div>
      </header>
    </>
  );
};

export default Header;
