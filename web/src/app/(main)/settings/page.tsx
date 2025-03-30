"use client";
import { AlertCircleIcon, CopyIcon, RefreshIcon } from "@/components/icons";
import { IFileType } from "@/models/file_type";
import { IFormResult } from "@/models/form";
import {
  IFormValidationErrors,
  IProjectFormProps,
  Project,
} from "@/models/project";
import { useAxios } from "@/providers/axios";
import { useActiveProject } from "@/providers/project";
import { useSession } from "next-auth/react";
import Link from "next/link";
import React, { FormEvent, useEffect, useState } from "react";
import { ScaleLoader } from "react-spinners";

const ProjectSettingsPage: React.FC = () => {
  const { activeProject } = useActiveProject();
  const { data: session } = useSession();

  const [isLoading, setIsLoading] = useState(false);
  const [regeneratingKey, setRegeneratingKey] = useState(false);
  const [formResult, setFormResult] = useState<IFormResult>({
    success: false,
    message: "",
    showResult: false,
  });
  const [projectKey, setProjectKey] = useState("proj_5f8a7b3c9d2e1f0a4b6c8d7e");

  const [formData, setFormData] = useState<IProjectFormProps>({
    projectName: "",
    description: "",
    maxFileSize: 5,
    allowedFiles: ["image/jpeg", "image/png", "application/pdf"],
  });

  const [errors, setErrors] = useState<IFormValidationErrors>({
    projectName: "",
    maxFileSize: "",
  });
  const [mimeTypes, setMimeTypes] = useState<IFileType[]>([]);
  const { get, post, put } = useAxios();
  const { setActiveProject } = useActiveProject();

  useEffect(() => {
    async function fetchFileTypes() {
      if (session && session.user) {
        const response = await get<IFileType[]>("/file-types");
        if (response.success) {
          setMimeTypes(response.result);
        } else {
          console.error("Failed to fetch file types");
        }
      }
    }

    async function fetchProjectDetails() {
      if (activeProject && session && session.user) {
        const response = await get<Project>(
          `/projects/${activeProject.id}/project-info`
        );
        if (response.success) {
          setFormData({
            projectName: response.result.name,
            description: response.result.description,
            maxFileSize: response.result.max_upload_size,
            allowedFiles: response.result.allowed_file_types,
          });
          setProjectKey(response.result.project_key);
        } else {
          console.error("Failed to fetch project details");
        }
      }
    }

    fetchFileTypes();
    fetchProjectDetails();
  }, [activeProject, get, session]);

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

  const handleSubmit = async (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    setFormResult({
      success: true,
      message: "",
      showResult: false,
    });

    if (validateForm()) {
      setIsLoading(true);

      const payload = {
        id: activeProject!.id,
        name: formData.projectName,
        description: formData.description,
        max_upload_size: parseInt(formData.maxFileSize.toString()),
        allowed_file_types: formData.allowedFiles,
      };

      try {
        const response = await put<Project>("/projects", payload);
        setIsLoading(false);
        setFormResult({
          success: true,
          message: "Project settings updated successfully",
          showResult: true,
        });
        setActiveProject(response.result);
      } catch (e) {
        console.log("Error: ", e);
        setIsLoading(false);
        setFormResult({
          success: false,
          message: "Failed to update project settings",
          showResult: true,
        });
      }

      // Clear errors / messages after 1 second
      setTimeout(() => {
        setFormResult({
          success: true,
          message: "",
          showResult: false,
        });
      }, 4000);
    }
  };

  const handleRegenerateKey = async () => {
    setRegeneratingKey(true);

    try {
      const payload = {
        id: activeProject?.id,
      };

      const response = await post<Project>(
        "/projects/re-generate-key",
        payload
      );
      setFormResult({
        success: true,
        message: "API key regenerated successfully",
        showResult: true,
      });

      setProjectKey(response.result.project_key);
      setRegeneratingKey(false);
    } catch (e) {
      console.log("Error: ", e);
      setRegeneratingKey(false);
      setFormResult({
        success: false,
        message: "Failed to regenerate API key",
        showResult: true,
      });
    }

    // Clear messages after 3 seconds
    setTimeout(() => {
      setFormResult({
        success: false,
        message: "",
        showResult: false,
      });
    }, 4000);
  };

  const handleProjectDeletion = (e: React.MouseEvent<HTMLButtonElement>) => {
    e.preventDefault();
    console.log("Project deletion initiated");
    const confirmDelete = confirm(
      "Are you sure you want to delete this project? This will delete all the files uploaded to this project, and cannot be undone."
    );
    if (confirmDelete) {
      // Delete project
      // TODO: Handle project deletion
    }
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

        {formResult.showResult && (
          <div
            className={`mb-6 p-3 border rounded-md text-sm ${
              formResult.success
                ? "bg-green-50 border-green-200 text-green-700"
                : "bg-red-50 border-red-200 text-red-700"
            }`}
          >
            {formResult.message}
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
                  } rounded-md focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm`}
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
                  className="shadow-sm block w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm"
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
                  } rounded-md focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm`}
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
                  className="shadow-sm block w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm"
                  size={5}
                >
                  {mimeTypes.map((fileType) => (
                    <option key={fileType.id} value={fileType.mimetype}>
                      {fileType.name}
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
            <h2 className="text-lg font-medium text-gray-900">Project Key</h2>
            <p className="mt-1 text-sm text-gray-500">
              This key is used to authenticate requests to the API for file
              uploads. Keep it secure.
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

                    setTimeout(() => {}, 3000);
                  }}
                  className="inline-flex items-center px-3 py-2 border border-gray-300 shadow-sm text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 mr-2 cursor-pointer"
                >
                  {/* Copy icon as SVG */}
                  <CopyIcon className="mr-2 h-4 w-4" />
                  Copy
                </button>
                <button
                  type="button"
                  onClick={handleRegenerateKey}
                  className="inline-flex items-center px-3 py-2 border border-gray-300 shadow-sm text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 cursor-pointer"
                  disabled={regeneratingKey}
                >
                  {regeneratingKey ? (
                    <ScaleLoader loading={true} color="#6366F1" height={15} />
                  ) : (
                    <>
                      {/* Refresh/Regenerate icon as SVG */}
                      <RefreshIcon className="mr-2 h-4 w-4" />
                      Regenerate
                    </>
                  )}
                </button>
              </div>
            </div>
            <div className="mt-2 flex items-start">
              <div className="flex-shrink-0">
                {/* Alert Circle icon as SVG */}
                <AlertCircleIcon className="h-5 w-5 text-yellow-400" />
              </div>
              <div className="ml-2">
                <p className="text-sm text-yellow-700">
                  Regenerating your Project key will invalidate the old key
                  immediately. Make sure to update any systems using this key.
                </p>
              </div>
            </div>
          </div>

          <div className="mt-8 flex space-x-4">
            <button
              type="submit"
              className="w-1/4 justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 cursor-pointer"
              disabled={isLoading}
            >
              {isLoading ? (
                <ScaleLoader loading={true} color="#fff" height={20} />
              ) : (
                "Save Changes"
              )}
            </button>
            <button
              type="button"
              className="w-1/4 justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-red-600 hover:bg-red-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500 cursor-pointer"
              disabled={isLoading}
              onClick={handleProjectDeletion}
            >
              {isLoading ? (
                <ScaleLoader loading={true} color="#fff" height={20} />
              ) : (
                "Delete Project"
              )}
            </button>
          </div>
        </form>
      </div>
    </div>
  );
};

export default ProjectSettingsPage;
