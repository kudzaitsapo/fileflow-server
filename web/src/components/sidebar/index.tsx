"use client";
import { Project } from "@/models/project";
import { useAxios } from "@/providers/axios";
import React, { useEffect } from "react";
import SidebarItem from "../sidebar-item";
import { signOut, useSession } from "next-auth/react";
import { ApiError } from "@/models/error";
import { useActiveProject } from "@/providers/project";

const SideBar: React.FC = () => {
  const { get } = useAxios();
  const { data: session } = useSession();
  const [projects, setProjects] = React.useState<Project[]>([]);
  const [activeIndex, setActiveIndex] = React.useState<number>(0);
  const { setActiveProject } = useActiveProject();

  useEffect(() => {
    const fetchProjects = async () => {
      if (session && session.user) {
        const response = await get<Project[] | ApiError>("/projects");
        if (Array.isArray(response)) {
          setProjects(response);
          setActiveProject(response[0]);
        }
      }
    };

    fetchProjects();
  }, [get, session, setActiveProject]);

  const setCurrentlyActiveProject = (index: number) => {
    setActiveIndex(index);
    setActiveProject(projects[index]);
  };

  const logOut = async () => {
    await signOut();
  };

  return (
    <>
      <div className="sidebar">
        <button className="create-btn">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            width="16"
            height="16"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            strokeWidth="2"
            strokeLinecap="round"
            strokeLinejoin="round"
          >
            <line x1="12" y1="5" x2="12" y2="19"></line>
            <line x1="5" y1="12" x2="19" y2="12"></line>
          </svg>
          New Project
        </button>

        <div className="sidebar-section">
          <div className="sidebar-title">Your Projects</div>
          <div className="project-list">
            {projects &&
              projects.map((project, index) => (
                <SidebarItem
                  name={project.name}
                  onClick={() => setCurrentlyActiveProject(index)}
                  isActive={activeIndex == index}
                  key={project.id}
                />
              ))}
          </div>
        </div>

        <div className="sidebar-section">
          <div className="sidebar-title">Settings</div>
          <div className="sidebar-item">
            <svg
              className="sidebar-icon"
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 24 24"
              fill="none"
              stroke="currentColor"
              strokeWidth="2"
              strokeLinecap="round"
              strokeLinejoin="round"
            >
              <path d="M21 15a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2z"></path>
            </svg>
            Activity Log
          </div>
          <div className="sidebar-item">
            <svg
              className="sidebar-icon"
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 24 24"
              fill="none"
              stroke="currentColor"
              strokeWidth="2"
              strokeLinecap="round"
              strokeLinejoin="round"
            >
              <line x1="12" y1="20" x2="12" y2="10"></line>
              <line x1="18" y1="20" x2="18" y2="4"></line>
              <line x1="6" y1="20" x2="6" y2="16"></line>
            </svg>
            Usage & Analytics
          </div>
          <div className="sidebar-item">
            <svg
              className="sidebar-icon"
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 24 24"
              fill="none"
              stroke="currentColor"
              strokeWidth="2"
              strokeLinecap="round"
              strokeLinejoin="round"
            >
              <circle cx="12" cy="12" r="3"></circle>
              <path d="M19.4 15a1.65 1.65 0 0 0 .33 1.82l.06.06a2 2 0 0 1 0 2.83 2 2 0 0 1-2.83 0l-.06-.06a1.65 1.65 0 0 0-1.82-.33 1.65 1.65 0 0 0-1 1.51V21a2 2 0 0 1-2 2 2 2 0 0 1-2-2v-.09A1.65 1.65 0 0 0 9 19.4a1.65 1.65 0 0 0-1.82.33l-.06.06a2 2 0 0 1-2.83 0 2 2 0 0 1 0-2.83l.06-.06a1.65 1.65 0 0 0 .33-1.82 1.65 1.65 0 0 0-1.51-1H3a2 2 0 0 1-2-2 2 2 0 0 1 2-2h.09A1.65 1.65 0 0 0 4.6 9a1.65 1.65 0 0 0-.33-1.82l-.06-.06a2 2 0 0 1 0-2.83 2 2 0 0 1 2.83 0l.06.06a1.65 1.65 0 0 0 1.82.33H9a1.65 1.65 0 0 0 1-1.51V3a2 2 0 0 1 2-2 2 2 0 0 1 2 2v.09a1.65 1.65 0 0 0 1 1.51 1.65 1.65 0 0 0 1.82-.33l.06-.06a2 2 0 0 1 2.83 0 2 2 0 0 1 0 2.83l-.06.06a1.65 1.65 0 0 0-.33 1.82V9a1.65 1.65 0 0 0 1.51 1H21a2 2 0 0 1 2 2 2 2 0 0 1-2 2h-.09a1.65 1.65 0 0 0-1.51 1z"></path>
            </svg>
            Account Settings
          </div>
          <div className="sidebar-item">
            <svg
              className="sidebar-icon"
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 24 24"
              fill="none"
              stroke="currentColor"
              strokeWidth="2"
              strokeLinecap="round"
              strokeLinejoin="round"
            >
              <path d="M12 22s8-4 8-10V5l-8-3-8 3v7c0 6 8 10 8 10z"></path>
            </svg>
            Security
          </div>
          <div className="sidebar-item">
            <svg
              className="sidebar-icon"
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 24 24"
              fill="none"
              stroke="currentColor"
              strokeWidth="2"
              strokeLinecap="round"
              strokeLinejoin="round"
            >
              <circle cx="12" cy="12" r="10"></circle>
              <line x1="12" y1="8" x2="12" y2="12"></line>
              <line x1="12" y1="16" x2="12.01" y2="16"></line>
            </svg>
            Help & Support
          </div>
        </div>

        <div className="sidebar-section">
          <div className="sidebar-title">Other</div>
          <div className="sidebar-item" onClick={logOut}>
            <svg
              className="sidebar-icon"
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 24 24"
              fill="none"
              stroke="currentColor"
              strokeWidth="2"
              strokeLinecap="round"
              strokeLinejoin="round"
            >
              <path d="M9 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h4"></path>
              <polyline points="16 17 21 12 16 7"></polyline>
              <line x1="21" y1="12" x2="9" y2="12"></line>
            </svg>
            Sign Out
          </div>
        </div>
      </div>
    </>
  );
};

export default SideBar;
