import { createContext } from "react";
import { dir } from "../../wailsjs/go/models";

interface PathData {
  dirList: dir.Dir[];
  paths: string[];
  selectedPaths: dir.Dir[];
  setDirList: (dirs: dir.Dir[]) => void;
  setPaths: (paths: string[]) => void;
  setSelectedPaths: (paths: dir.Dir[]) => void;
  handlePath: (path: string, dir: string) => void;
  getDirs: (path: string) => void;
}

export const PathContext = createContext<PathData>({
  dirList: [],
  paths: [],
  selectedPaths: [],
  setDirList: () => {},
  setPaths: () => {},
  setSelectedPaths: () => {},
  handlePath: () => {},
  getDirs: () => {},
});
