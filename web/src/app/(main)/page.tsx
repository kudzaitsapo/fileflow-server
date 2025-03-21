"use client";

import {
  ListIcon,
  SearchIcon,
  SettingsIcon,
  UploadIcon,
} from "@/components/icons";
import { StoredFile } from "@/models/file";
import { useAxios } from "@/providers/axios";
import { useActiveProject } from "@/providers/project";
import { useRouter } from "next/navigation";
import { useEffect, useState } from "react";

export default function Home() {
  const { activeProject } = useActiveProject();
  const [files, setFiles] = useState<StoredFile[]>([]);
  const { get } = useAxios();
  const router = useRouter();

  useEffect(() => {
    async function getProjectFiles(projectId: number) {
      const response = await get<StoredFile[]>(`/projects/${projectId}/files`, {
        limit: "10",
      });
      if (Array.isArray(response)) {
        setFiles(response);
      }
    }

    if (activeProject) {
      getProjectFiles(activeProject.id);
    }
  }, [get, activeProject]);

  const handleNavigation = () => {
    router.push("/projects/settings");
  };

  return (
    <div className="flex-1 p-6 overflow-y-auto">
      <div className="flex items-center gap-2 mb-6">
        <span className="text-gray-600 text-sm">Projects</span>
        <span className="text-gray-400">/</span>
        <span className="text-gray-800 text-sm font-medium">
          {activeProject?.name}
        </span>
      </div>

      <div className="flex items-center justify-between mb-6">
        <h1 className="text-2xl font-semibold">{activeProject?.name}</h1>
      </div>

      <div className="flex gap-3 mb-6">
        <div className="flex items-center bg-white border border-gray-300 rounded-md py-2 px-3 flex-1 max-w-md">
          <SearchIcon className="text-gray-500 mr-2" />
          <input
            type="text"
            className="border-none outline-none flex-1 text-sm"
            placeholder="Search files..."
          />
        </div>

        <button className="flex items-center gap-2 py-2 px-3 rounded-md bg-white border border-gray-300 text-gray-700 font-medium text-sm cursor-pointer hover:bg-gray-50">
          <ListIcon className="w-4 h-4" />
          Filter
        </button>

        <button className="flex items-center gap-2 py-2 px-3 rounded-md bg-blue-600 text-white border-none font-medium text-sm cursor-pointer hover:bg-blue-800">
          <UploadIcon className="w-4 h-4" />
          Upload Files
        </button>

        <button
          className="flex items-center gap-2 py-2 px-3 rounded-md bg-white border border-gray-300 text-gray-700 font-medium text-sm cursor-pointer hover:bg-gray-50"
          onClick={() => handleNavigation()}
        >
          <SettingsIcon className="w-4 h-4" />
          Project Settings
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
            {!files.length && (
              <tr className="h-64">
                <td colSpan={5} className="text-center">
                  <div className="flex flex-col items-center justify-center p-8">
                    <div className="mb-4 text-gray-300">
                      <svg
                        xmlns="http://www.w3.org/2000/svg"
                        width="48"
                        height="48"
                        viewBox="0 0 24 24"
                        fill="none"
                        stroke="currentColor"
                        strokeWidth="1.5"
                        strokeLinecap="round"
                        strokeLinejoin="round"
                      >
                        <rect
                          x="3"
                          y="3"
                          width="18"
                          height="18"
                          rx="2"
                          ry="2"
                        ></rect>
                        <line x1="3" y1="9" x2="21" y2="9"></line>
                        <line x1="9" y1="21" x2="9" y2="9"></line>
                      </svg>
                    </div>
                    <h3 className="mb-2 text-xl text-gray-600">
                      No files found
                    </h3>
                    <p className="text-sm text-gray-500">
                      There are no files to display at this time.
                    </p>
                  </div>
                </td>
              </tr>
            )}

            {files &&
              files.length > 0 &&
              files.map((file) => (
                <tr key={file.id} className="hover:bg-gray-50">
                  <td className="py-3 px-4 border-b border-gray-200">
                    <div className="flex items-center gap-3">
                      <div className="w-8 h-8 flex items-center justify-center rounded bg-red-50 text-red-500">
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
                        <div className="text-xs text-gray-500">
                          /web-app/docs/
                        </div>
                      </div>
                    </div>
                  </td>
                  <td className="py-3 px-4 border-b border-gray-200 text-sm text-gray-600">
                    3.2 MB
                  </td>
                  <td className="py-3 px-4 border-b border-gray-200 text-sm">
                    PDF Document
                  </td>
                  <td className="py-3 px-4 border-b border-gray-200 text-sm text-gray-600">
                    March 4, 2025
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
}
