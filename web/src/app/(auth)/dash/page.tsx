"use client";
import React, { useState } from "react";

// Define TypeScript interfaces
interface FileItem {
  name: string;
  path: string;
  size: string;
  type: string;
  date: string;
  icon: "pdf" | "doc" | "xls" | "img" | "zip" | "txt" | "video" | "audio";
}

interface Project {
  name: string;
}

// Header Component
const Header: React.FC = () => {
  return (
    <header className="bg-white shadow sticky top-0 z-10 flex items-center justify-between px-6 py-3">
      <div className="flex items-center gap-2 font-bold text-xl text-blue-800">
        <span className="w-6 h-6 bg-blue-600 rounded flex items-center justify-center text-white font-bold">
          F
        </span>
        <span>FileFlow</span>
      </div>
      <div className="flex items-center gap-4">
        <div className="w-8 h-8 rounded-full bg-gray-300 flex items-center justify-center text-gray-700 font-semibold">
          JS
        </div>
      </div>
    </header>
  );
};

// Sidebar Component
interface SidebarProps {
  activeProject: string;
  projects: Project[];
  onSelectProject: (project: string) => void;
}

const Sidebar: React.FC<SidebarProps> = ({
  activeProject,
  projects,
  onSelectProject,
}) => {
  return (
    <div className="w-64 bg-white border-r border-gray-200 p-4 pt-6 flex flex-col gap-6">
      <button className="bg-blue-600 text-white border-none rounded-md py-3 px-4 font-semibold cursor-pointer flex items-center justify-center gap-2 transition-colors hover:bg-blue-800">
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

      <div className="flex flex-col gap-2">
        <div className="text-xs font-semibold uppercase text-gray-500 px-2 mb-1">
          General
        </div>
        <div
          className={`flex items-center gap-3 p-2 rounded-md text-sm cursor-pointer ${
            activeProject === "Projects"
              ? "bg-blue-500 text-white"
              : "text-gray-600 hover:bg-gray-100"
          }`}
          onClick={() => onSelectProject("Projects")}
        >
          <svg
            className={`w-4 h-4 ${
              activeProject === "Projects" ? "text-white" : "text-gray-500"
            }`}
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            strokeWidth="2"
            strokeLinecap="round"
            strokeLinejoin="round"
          >
            <rect x="3" y="3" width="18" height="18" rx="2" ry="2"></rect>
            <line x1="3" y1="9" x2="21" y2="9"></line>
            <line x1="9" y1="21" x2="9" y2="9"></line>
          </svg>
          Projects
        </div>
        <div className="flex items-center gap-3 p-2 rounded-md text-gray-600 text-sm cursor-pointer hover:bg-gray-100">
          <svg
            className="w-4 h-4 text-gray-500"
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
        <div className="flex items-center gap-3 p-2 rounded-md text-gray-600 text-sm cursor-pointer hover:bg-gray-100">
          <svg
            className="w-4 h-4 text-gray-500"
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
      </div>

      <div className="flex flex-col gap-2">
        <div className="text-xs font-semibold uppercase text-gray-500 px-2 mb-1">
          Your Projects
        </div>
        <div className="overflow-y-auto flex-1">
          {projects.map((project, index) => (
            <div
              key={index}
              className={`flex items-center gap-3 p-2 rounded-md text-sm cursor-pointer ${
                activeProject === project.name
                  ? "bg-blue-500 text-white"
                  : "text-gray-600 hover:bg-gray-100"
              }`}
              onClick={() => onSelectProject(project.name)}
            >
              <svg
                className={`w-4 h-4 ${
                  activeProject === project.name
                    ? "text-white"
                    : "text-gray-500"
                }`}
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
              {project.name}
            </div>
          ))}
        </div>
      </div>

      <div className="flex flex-col gap-2">
        <div className="text-xs font-semibold uppercase text-gray-500 px-2 mb-1">
          Settings
        </div>
        <div className="flex items-center gap-3 p-2 rounded-md text-gray-600 text-sm cursor-pointer hover:bg-gray-100">
          <svg
            className="w-4 h-4 text-gray-500"
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
        <div className="flex items-center gap-3 p-2 rounded-md text-gray-600 text-sm cursor-pointer hover:bg-gray-100">
          <svg
            className="w-4 h-4 text-gray-500"
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
        <div className="flex items-center gap-3 p-2 rounded-md text-gray-600 text-sm cursor-pointer hover:bg-gray-100">
          <svg
            className="w-4 h-4 text-gray-500"
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
    </div>
  );
};

// Main Content Component
interface MainContentProps {
  activeProject: string;
  files: FileItem[];
}

const MainContent: React.FC<MainContentProps> = ({ activeProject, files }) => {
  return (
    <div className="flex-1 p-6 overflow-y-auto">
      <div className="flex items-center gap-2 mb-6">
        <span className="text-gray-600 text-sm">Projects</span>
        <span className="text-gray-400">/</span>
        <span className="text-gray-800 text-sm font-medium">
          {activeProject}
        </span>
      </div>

      <div className="flex items-center justify-between mb-6">
        <h1 className="text-2xl font-semibold">{activeProject} Project</h1>
      </div>

      <div className="flex gap-3 mb-6">
        <div className="flex items-center bg-white border border-gray-300 rounded-md py-2 px-3 flex-1 max-w-md">
          <svg
            className="text-gray-500 mr-2"
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
            <circle cx="11" cy="11" r="8"></circle>
            <line x1="21" y1="21" x2="16.65" y2="16.65"></line>
          </svg>
          <input
            type="text"
            className="border-none outline-none flex-1 text-sm"
            placeholder="Search files..."
          />
        </div>

        <button className="flex items-center gap-2 py-2 px-3 rounded-md bg-white border border-gray-300 text-gray-700 font-medium text-sm cursor-pointer hover:bg-gray-50">
          <svg
            className="w-4 h-4"
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            strokeWidth="2"
            strokeLinecap="round"
            strokeLinejoin="round"
          >
            <line x1="21" y1="10" x2="3" y2="10"></line>
            <line x1="21" y1="6" x2="3" y2="6"></line>
            <line x1="21" y1="14" x2="3" y2="14"></line>
            <line x1="21" y1="18" x2="3" y2="18"></line>
          </svg>
          Filter
        </button>

        <button className="flex items-center gap-2 py-2 px-3 rounded-md bg-blue-600 text-white border-none font-medium text-sm cursor-pointer hover:bg-blue-800">
          <svg
            className="w-4 h-4"
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            strokeWidth="2"
            strokeLinecap="round"
            strokeLinejoin="round"
          >
            <path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4"></path>
            <polyline points="17 8 12 3 7 8"></polyline>
            <line x1="12" y1="3" x2="12" y2="15"></line>
          </svg>
          Upload Files
        </button>

        <button className="flex items-center gap-2 py-2 px-3 rounded-md bg-white border border-gray-300 text-gray-700 font-medium text-sm cursor-pointer hover:bg-gray-50">
          <svg
            className="w-4 h-4"
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            strokeWidth="2"
            strokeLinecap="round"
            strokeLinejoin="round"
          >
            <polyline points="3 6 5 6 21 6"></polyline>
            <path d="M19 6v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6m3 0V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2"></path>
          </svg>
          New Folder
        </button>
      </div>

      <div className="bg-white rounded-lg border border-gray-200 overflow-hidden">
        <table className="w-full border-collapse">
          <thead>
            <tr>
              <th className="text-left py-3 px-4 text-xs uppercase text-gray-500 border-b border-gray-200 bg-gray-50">
                File Name
              </th>
              <th className="text-left py-3 px-4 text-xs uppercase text-gray-500 border-b border-gray-200 bg-gray-50">
                Size
              </th>
              <th className="text-left py-3 px-4 text-xs uppercase text-gray-500 border-b border-gray-200 bg-gray-50">
                File Type
              </th>
              <th className="text-left py-3 px-4 text-xs uppercase text-gray-500 border-b border-gray-200 bg-gray-50">
                Date Uploaded
              </th>
              <th className="text-left py-3 px-4 text-xs uppercase text-gray-500 border-b border-gray-200 bg-gray-50">
                Actions
              </th>
            </tr>
          </thead>
          <tbody>
            {files.map((file, index) => (
              <tr key={index} className="hover:bg-gray-50">
                <td className="py-3 px-4 border-b border-gray-200">
                  <div className="flex items-center gap-3">
                    <div
                      className={`w-8 h-8 flex items-center justify-center rounded bg-red-50 text-red-500 ${
                        file.icon === "pdf"
                          ? "bg-red-50 text-red-500"
                          : file.icon === "xls"
                          ? "bg-green-50 text-green-500"
                          : "bg-gray-50 text-gray-500"
                      }`}
                    >
                      <svg
                        xmlns="http://www.w3.org/2000/svg"
                        width="20"
                        height="20"
                        viewBox="0 0 24 24"
                        fill="none"
                        stroke="currentColor"
                        strokeWidth="2"
                        strokeLinecap="round"
                        strokeLinejoin="round"
                      >
                        <path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"></path>
                        <polyline points="14 2 14 8 20 8"></polyline>
                        <path d="M9 15h6"></path>
                        <path d="M9 11h6"></path>
                      </svg>
                    </div>
                    <div>
                      <div className="text-sm font-medium text-gray-800">
                        {file.name}
                      </div>
                      <div className="text-xs text-gray-500">{file.path}</div>
                    </div>
                  </div>
                </td>
                <td className="py-3 px-4 border-b border-gray-200 text-sm text-gray-600">
                  {file.size}
                </td>
                <td className="py-3 px-4 border-b border-gray-200 text-sm">
                  {file.type}
                </td>
                <td className="py-3 px-4 border-b border-gray-200 text-sm text-gray-600">
                  {file.date}
                </td>
                <td className="py-3 px-4 border-b border-gray-200">
                  <div className="flex gap-1">
                    <button className="bg-transparent border-none cursor-pointer p-1.5 rounded text-gray-500 hover:bg-gray-100 hover:text-gray-900">
                      <svg
                        xmlns="http://www.w3.org/2000/svg"
                        width="18"
                        height="18"
                        viewBox="0 0 24 24"
                        fill="none"
                        stroke="currentColor"
                        strokeWidth="2"
                        strokeLinecap="round"
                        strokeLinejoin="round"
                      >
                        <path d="M1 12s4-8 11-8 11 8 11 8-4 8-11 8-11-8-11-8z"></path>
                        <circle cx="12" cy="12" r="3"></circle>
                      </svg>
                    </button>
                    <button className="bg-transparent border-none cursor-pointer p-1.5 rounded text-gray-500 hover:bg-gray-100 hover:text-gray-900">
                      <svg
                        xmlns="http://www.w3.org/2000/svg"
                        width="18"
                        height="18"
                        viewBox="0 0 24 24"
                        fill="none"
                        stroke="currentColor"
                        strokeWidth="2"
                        strokeLinecap="round"
                        strokeLinejoin="round"
                      >
                        <path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4"></path>
                        <polyline points="7 10 12 15 17 10"></polyline>
                        <line x1="12" y1="15" x2="12" y2="3"></line>
                      </svg>
                    </button>
                    <button className="bg-transparent border-none cursor-pointer p-1.5 rounded text-gray-500 hover:bg-gray-100 hover:text-gray-900">
                      <svg
                        xmlns="http://www.w3.org/2000/svg"
                        width="18"
                        height="18"
                        viewBox="0 0 24 24"
                        fill="none"
                        stroke="currentColor"
                        strokeWidth="2"
                        strokeLinecap="round"
                        strokeLinejoin="round"
                      >
                        <path d="M3 6h18"></path>
                        <path d="M19 6v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6m3 0V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2"></path>
                      </svg>
                    </button>
                  </div>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
};

// Main App Component
const FileServerApp: React.FC = () => {
  const [activeProject, setActiveProject] = useState<string>("Web App");

  const projects: Project[] = [
    { name: "Web App" },
    { name: "Mobile App" },
    { name: "Machine Learning" },
    { name: "API Integration" },
  ];

  const files: FileItem[] = [
    {
      name: "Technical_Specification.pdf",
      path: "/web-app/docs/",
      size: "3.2 MB",
      type: "PDF Document",
      date: "March 4, 2025",
      icon: "pdf",
    },
    {
      name: "Project_Budget.xlsx",
      path: "/finance",
      size: "1.5 MB",
      type: "Excel Spreadsheet",
      date: "March 6, 2025",
      icon: "xls",
    },
  ];

  const handleSelectProject = (project: string) => {
    setActiveProject(project);
  };

  return (
    <div className="bg-gray-100 text-gray-800">
      <Header />
      <div className="flex h-[calc(100vh-4rem)]">
        <Sidebar
          activeProject={activeProject}
          projects={projects}
          onSelectProject={handleSelectProject}
        />
        <MainContent activeProject={activeProject} files={files} />
      </div>
    </div>
  );
};

export default FileServerApp;
