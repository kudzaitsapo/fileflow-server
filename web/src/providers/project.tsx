"use client";
import React, {
  createContext,
  ReactNode,
  useContext,
  useEffect,
  useState,
} from "react";
import { Project } from "@/models/project";
import { ACTIVE_PROJECT_LOCAL_STORAGE_KEY } from "@/constants/app";
import { useCookiesNext } from "cookies-next";

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
  const [localActiveProject, setLocalActiveProject] = useState<Project | null>(
    null
  );
  const { getCookie, setCookie, deleteCookie } = useCookiesNext();

  useEffect(() => {
    const storedProject = getCookie(ACTIVE_PROJECT_LOCAL_STORAGE_KEY);
    if (storedProject) {
      const parsedProject = JSON.parse(storedProject) as Project;
      setLocalActiveProject(parsedProject);
    } else {
      setLocalActiveProject(null);
    }
  }, [getCookie]);

  const setActiveProject = (project: Project | null) => {
    if (project) {
      setLocalActiveProject(project);
      setCookie(ACTIVE_PROJECT_LOCAL_STORAGE_KEY, JSON.stringify(project), {
        maxAge: 60 * 60 * 24 * 30, // 30 days
      });
    } else {
      deleteCookie(ACTIVE_PROJECT_LOCAL_STORAGE_KEY);
      setLocalActiveProject(null);
    }
  };

  return (
    <ActiveProjectContext.Provider
      value={{
        activeProject: localActiveProject,
        setActiveProject,
      }}
    >
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
