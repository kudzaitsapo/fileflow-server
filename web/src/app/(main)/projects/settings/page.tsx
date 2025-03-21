"use client";
import { IFormValidationErrors, IProjectFormProps } from "@/models/project";
import { useActiveProject } from "@/providers/project";
import Link from "next/link";
//import { useParams } from "next/navigation";
import React, { useEffect, useState } from "react";
import { ScaleLoader } from "react-spinners";

const ProjectSettingsPage: React.FC = () => {
  //const { id } = useParams<{ id: string }>();

  const { activeProject } = useActiveProject();

  const [isLoading, setIsLoading] = useState(false);
  const [regeneratingKey, setRegeneratingKey] = useState(false);
  const [successMessage, setSuccessMessage] = useState("");
  const [projectKey, setProjectKey] = useState("proj_5f8a7b3c9d2e1f0a4b6c8d7e");

  const [formData, setFormData] = useState<IProjectFormProps>({
    projectName: "",
    description: "",
    maxFileSize: 5,
    allowedFiles: ["image/jpeg", "image/png", "application/pdf"],
  });

  console.log("LOG::activeProject: ", activeProject);

  const [errors, setErrors] = useState<IFormValidationErrors>({
    projectName: "",
    maxFileSize: "",
  });

  const mimeTypeOptions = [
    { value: "image/jpeg", label: "JPEG Images (.jpg, .jpeg)" },
    { value: "image/png", label: "PNG Images (.png)" },
    { value: "image/gif", label: "GIF Images (.gif)" },
    { value: "image/svg+xml", label: "SVG Images (.svg)" },
    { value: "application/pdf", label: "PDF Documents (.pdf)" },
    { value: "application/msword", label: "Word Documents (.doc)" },
    {
      value:
        "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
      label: "Word Documents (.docx)",
    },
    { value: "text/plain", label: "Text Files (.txt)" },
    { value: "text/csv", label: "CSV Files (.csv)" },
    { value: "application/zip", label: "ZIP Archives (.zip)" },
  ];

  useEffect(() => {
    if (activeProject) {
      setFormData({
        projectName: activeProject.name,
        description: activeProject.description,
        maxFileSize: activeProject.max_file_size,
        allowedFiles: ["image/jpeg", "image/png", "application/pdf"],
      });
    }
    console.log("LOG::activeProject: ", activeProject);
  }, [activeProject]);

  const handleInputChange = (
    e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>
  ) => {
    const { name, value } = e.target;
    setFormData({
      ...formData,
      [name]: value,
    });

    // Clear error when user types
    if (errors[name as keyof typeof errors]) {
      setErrors({
        ...errors,
        [name]: "",
      });
    }
  };

  const handleMimeTypeChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
    const selectedOptions = Array.from(e.target.selectedOptions).map(
      (option: HTMLOptionElement) => option.value
    );
    setFormData({
      ...formData,
      allowedFiles: selectedOptions,
    });
  };

  const validateForm = () => {
    const newErrors: IFormValidationErrors = {};

    if (!formData?.projectName?.trim()) {
      newErrors.projectName = "Project name is required";
    }

    if (!formData.maxFileSize || formData.maxFileSize <= 0) {
      newErrors.maxFileSize = "Max file size must be greater than 0";
    }

    setErrors(newErrors);
    return Object.keys(newErrors).length === 0;
  };

  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  const handleSubmit = (e: any) => {
    e.preventDefault();

    if (validateForm()) {
      setIsLoading(true);

      // Simulate API call
      setTimeout(() => {
        setIsLoading(false);
        setSuccessMessage("Project settings updated successfully!");

        // Clear success message after 3 seconds
        setTimeout(() => {
          setSuccessMessage("");
        }, 3000);
      }, 1000);
    }
  };

  const handleRegenerateKey = () => {
    setRegeneratingKey(true);

    // Simulate API call to regenerate key
    setTimeout(() => {
      const newKey = "proj_" + Math.random().toString(36).substr(2, 20);
      setProjectKey(newKey);
      setRegeneratingKey(false);
      setSuccessMessage("Project key regenerated successfully!");

      // Clear success message after 3 seconds
      setTimeout(() => {
        setSuccessMessage("");
      }, 3000);
    }, 1000);
  };

  return (
    <div className="flex-1 p-6 overflow-y-auto">
      <div className="flex items-center gap-2 mb-6">
        <Link href="/" className="text-gray-600 text-sm">
          Projects
        </Link>
        <span className="text-gray-400">/</span>
        <Link href={`/`} className="text-gray-600 text-sm">
          {activeProject?.name}
        </Link>
        <span className="text-gray-400">/</span>
        <span className="text-gray-800 text-sm font-medium">Settings</span>
      </div>

      <div className="bg-white rounded-lg border border-gray-200 overflow-hidden p-6">
        <div className="mb-6">
          <h1 className="text-2xl font-bold text-gray-900">Project Settings</h1>
          <p className="text-sm text-gray-500 mt-1">
            Manage your project configuration and security settings
          </p>
        </div>

        {successMessage && (
          <div className="mb-6 p-3 bg-green-50 border border-green-200 text-green-700 rounded-md text-sm">
            {successMessage}
          </div>
        )}

        <form onSubmit={handleSubmit}>
          <div className="space-y-6">
            {/* Project Name */}
            <div>
              <label
                htmlFor="projectName"
                className="block text-sm font-medium text-gray-700"
              >
                Project Name <span className="text-red-500">*</span>
              </label>
              <div className="mt-1">
                <input
                  type="text"
                  id="projectName"
                  name="projectName"
                  value={formData.projectName}
                  onChange={handleInputChange}
                  className={`shadow-sm block w-full px-3 py-2 border ${
                    errors.projectName ? "border-red-300" : "border-gray-300"
                  } rounded-md focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm`}
                />
                {errors.projectName && (
                  <p className="mt-1 text-sm text-red-600">
                    {errors.projectName}
                  </p>
                )}
              </div>
            </div>

            {/* Description */}
            <div>
              <label
                htmlFor="description"
                className="block text-sm font-medium text-gray-700"
              >
                Description
              </label>
              <div className="mt-1">
                <textarea
                  id="description"
                  name="description"
                  rows={3}
                  value={formData.description}
                  onChange={handleInputChange}
                  className="shadow-sm block w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
                  placeholder="Brief description of your project"
                />
              </div>
            </div>

            {/* Max File Size */}
            <div>
              <label
                htmlFor="maxFileSize"
                className="block text-sm font-medium text-gray-700"
              >
                Max File Size (MB) <span className="text-red-500">*</span>
              </label>
              <div className="mt-1">
                <input
                  type="number"
                  id="maxFileSize"
                  name="maxFileSize"
                  value={formData.maxFileSize}
                  onChange={handleInputChange}
                  min="1"
                  className={`shadow-sm block w-full px-3 py-2 border ${
                    errors.maxFileSize ? "border-red-300" : "border-gray-300"
                  } rounded-md focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm`}
                />
                {errors.maxFileSize && (
                  <p className="mt-1 text-sm text-red-600">
                    {errors.maxFileSize}
                  </p>
                )}
              </div>
            </div>

            {/* Allowed Files */}
            <div>
              <label
                htmlFor="allowedFiles"
                className="block text-sm font-medium text-gray-700"
              >
                Allowed File Types
              </label>
              <div className="mt-1">
                <select
                  id="allowedFiles"
                  name="allowedFiles"
                  multiple
                  value={formData.allowedFiles}
                  onChange={handleMimeTypeChange}
                  className="shadow-sm block w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
                  size={5}
                >
                  {mimeTypeOptions.map((option) => (
                    <option key={option.value} value={option.value}>
                      {option.label}
                    </option>
                  ))}
                </select>
                <p className="mt-1 text-xs text-gray-500">
                  Hold Ctrl/Cmd to select multiple types
                </p>
              </div>
            </div>
          </div>

          {/* Project Key Section */}
          <div className="mt-8 pt-6 border-t border-gray-200">
            <h2 className="text-lg font-medium text-gray-900">
              Project API Key
            </h2>
            <p className="mt-1 text-sm text-gray-500">
              This key is used to authenticate requests to your project&apos;s
              API
            </p>

            <div className="mt-4 flex items-center">
              <div className="flex-1 flex items-center">
                <input
                  type="text"
                  readOnly
                  value={projectKey}
                  className="block w-full px-3 py-2 bg-gray-50 border border-gray-300 rounded-md text-gray-500 sm:text-sm mr-2"
                />
                <button
                  type="button"
                  onClick={() => {
                    navigator.clipboard.writeText(projectKey);
                    setSuccessMessage("API key copied to clipboard!");
                    setTimeout(() => {
                      setSuccessMessage("");
                    }, 3000);
                  }}
                  className="inline-flex items-center px-3 py-2 border border-gray-300 shadow-sm text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 mr-2"
                >
                  {/* Copy icon as SVG */}
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    className="mr-2 h-4 w-4"
                    viewBox="0 0 24 24"
                    fill="none"
                    stroke="currentColor"
                    strokeWidth="2"
                    strokeLinecap="round"
                    strokeLinejoin="round"
                  >
                    <rect
                      x="9"
                      y="9"
                      width="13"
                      height="13"
                      rx="2"
                      ry="2"
                    ></rect>
                    <path d="M5 15H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v1"></path>
                  </svg>
                  Copy
                </button>
                <button
                  type="button"
                  onClick={handleRegenerateKey}
                  className="inline-flex items-center px-3 py-2 border border-gray-300 shadow-sm text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
                  disabled={regeneratingKey}
                >
                  {regeneratingKey ? (
                    <ScaleLoader loading={true} color="#6366F1" height={15} />
                  ) : (
                    <>
                      {/* Refresh/Regenerate icon as SVG */}
                      <svg
                        xmlns="http://www.w3.org/2000/svg"
                        className="mr-2 h-4 w-4"
                        viewBox="0 0 24 24"
                        fill="none"
                        stroke="currentColor"
                        strokeWidth="2"
                        strokeLinecap="round"
                        strokeLinejoin="round"
                      >
                        <path d="M21 2v6h-6"></path>
                        <path d="M3 12a9 9 0 0 1 15-6.7L21 8"></path>
                        <path d="M3 22v-6h6"></path>
                        <path d="M21 12a9 9 0 0 1-15 6.7L3 16"></path>
                      </svg>
                      Regenerate
                    </>
                  )}
                </button>
              </div>
            </div>
            <div className="mt-2 flex items-start">
              <div className="flex-shrink-0">
                {/* Alert Circle icon as SVG */}
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  className="h-5 w-5 text-yellow-400"
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
              </div>
              <div className="ml-2">
                <p className="text-sm text-yellow-700">
                  Regenerating your API key will invalidate the old key
                  immediately. Make sure to update any systems using this key.
                </p>
              </div>
            </div>
          </div>

          <div className="mt-8">
            <button
              type="submit"
              className="w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
              disabled={isLoading}
            >
              {isLoading ? (
                <ScaleLoader loading={true} color="#fff" height={20} />
              ) : (
                "Save Changes"
              )}
            </button>
          </div>
        </form>
      </div>
    </div>
  );
};

export default ProjectSettingsPage;
