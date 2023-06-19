import { createContext } from "react";
import { dir } from "../../wailsjs/go/models";

export interface PathData {
  dirList: dir.Dir[];
  paths: string[];
  selectedPaths: dir.Dir[];
  currentBody: string;
  setDirList: (dirs: dir.Dir[]) => void;
  setPaths: (paths: string[]) => void;
  setSelectedPaths: (paths: dir.Dir[]) => void;
  handlePath: (path: string, dir: string) => void;
  getDirs: (path: string) => void;
  setCurrentBody: (body: string) => void;
}

export const PathContext = createContext<PathData>({
  dirList: [],
  paths: [],
  selectedPaths: [],
  currentBody: "home",
  setDirList: () => {},
  setPaths: () => {},
  setSelectedPaths: () => {},
  handlePath: () => {},
  getDirs: () => {},
  setCurrentBody: () => {},
});
