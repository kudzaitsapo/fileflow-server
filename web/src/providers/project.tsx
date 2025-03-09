"use client";
import React, { createContext, ReactNode, useContext, useState } from "react";
import { Project } from "@/models/project";

interface ActiveProjectContextType {
  activeProject: Project | null;
  setActiveProject: (project: Project | null) => void;
}

const ActiveProjectContext = createContext<
  ActiveProjectContextType | undefined
>(undefined);

export const ActiveProjectProvider: React.FC<{ children: ReactNode }> = ({
  children,
}) => {
  const [activeProject, setActiveProject] = useState<Project | null>(null);

  return (
    <ActiveProjectContext.Provider value={{ activeProject, setActiveProject }}>
      {children}
    </ActiveProjectContext.Provider>
  );
};

export const useActiveProject = (): ActiveProjectContextType => {
  const context = useContext(ActiveProjectContext);
  if (!context) {
    throw new Error(
      "useActiveProject must be used within an ActiveProjectProvider"
    );
  }
  return context;
};
