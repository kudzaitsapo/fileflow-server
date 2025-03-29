"use client";

import { IFileType } from "@/models/file_type";
import {
  IFormValidationErrors,
  IProjectFormProps,
  Project,
} from "@/models/project";
import { useAxios } from "@/providers/axios";
import { useActiveProject } from "@/providers/project";
import { useSession } from "next-auth/react";
import Link from "next/link";
import React, { useEffect, FormEvent, useState } from "react";

import { ScaleLoader } from "react-spinners";

const ProjectCreationForm: React.FC = () => {
  const [formData, setFormData] = useState<IProjectFormProps>({
    projectName: "",
    description: "",
    allowedFiles: [],
    maxFileSize: 5,
  });

  const { post } = useAxios();
  const { data: session } = useSession();

  const [errors, setErrors] = useState<IFormValidationErrors>({});
  const [isLoading, setLoading] = useState<boolean>(false);
  const [successMessage, setSuccessMessage] = useState("");
  const { setActiveProject } = useActiveProject();
  const [mimeTypes, setMimeTypes] = useState<IFileType[]>([]);
  const { get } = useAxios();

  useEffect(() => {
    async function fetchFileTypes() {
      if (session && session.user) {
        const response = await get<IFileType[]>("/file-types");
        if (response.success) {
          setMimeTypes(response.result);
        } else {
          console.error("Error fetching file types");
        }
      }
    }

    fetchFileTypes();
  }, [get, session]);

  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  const handleInputChange = (e: any) => {
    const { name, value } = e.target;
    setFormData({
      ...formData,
      [name]: value,
    });
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

    if (!formData.projectName.trim()) {
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
    setLoading(true);

    if (validateForm()) {
      console.log("Form submitted:", formData);
      // TODO: allow user to select allowed file types
      const payload = {
        name: formData.projectName,
        description: formData.description,
        max_upload_size: parseInt(formData.maxFileSize.toString()),
      };
      try {
        const response = await post<Project>("/projects", payload);
        setLoading(false);
        setSuccessMessage("Successfully created project");
        setActiveProject(response.result);
      } catch (error) {
        console.error("LOG::error creating project: ", error);
        setLoading(false);
      }
    }
  };

  return (
    <div className="flex-1 p-6 overflow-y-auto">
      <div className="flex items-center gap-2 mb-6">
        <Link href="/" className="text-gray-600 text-sm">
          Projects
        </Link>
        <span className="text-gray-400">/</span>
        <span className="text-gray-800 text-sm font-medium">
          Create Project
        </span>
      </div>

      <div className="bg-white rounded-lg border border-gray-200 overflow-hidden p-6">
        <div className="mb-6">
          <h1 className="text-2xl font-bold text-gray-900">
            Create New Project
          </h1>
          <p className="text-sm text-gray-500 mt-1">
            Fill in the details below to create a new project
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
                  } rounded-md focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm`}
                />
                {errors.projectName && (
                  <p className="mt-1 text-sm text-red-600">
                    {errors.projectName}
                  </p>
                )}
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

          <div className="mt-8">
            <button
              type="submit"
              className="w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
              disabled={isLoading}
            >
              <ScaleLoader loading={isLoading} color="#fff" height={20} />
              Create Project
            </button>
          </div>
        </form>
      </div>
    </div>
  );
};

export default ProjectCreationForm;
