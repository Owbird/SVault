import {
  Box,
  Button,
  Checkbox,
  CheckboxGroup,
  Grid,
  GridItem,
  HStack,
  Text,
  VStack,
} from "@chakra-ui/react";
import { useEffect, useState } from "react";
import { dir } from "../wailsjs/go/models";
import {
  Encrypt,
  GetDirs,
  GetUserHome,
  OpenFile,
} from "../wailsjs/go/uifunctions/UIFunctions";

import { FcFile, FcFolder } from "react-icons/fc";
import SideBar from "./components/SideBar";
import { PathContext } from "./contexts/pathsContext";

const App = () => {
  const [dirList, setDirList] = useState<dir.Dir[]>([]);
  const [paths, setPaths] = useState<string[]>([]);
  const [selectedPaths, setSelectedPaths] = useState<dir.Dir[]>([]);

  const getDirs = (path: string) => {
    GetDirs(path).then((data) => {
      setDirList(data);
    });
    setSelectedPaths([]);
  };
  const handlePath = (path: string, dir: string) => {
    getDirs(path);
    setPaths([...paths, dir]);
  };

  useEffect(() => {
    GetUserHome().then((path) => handlePath(path, path));
  }, []);

  const handleSelected = () => {
    for (let dir of selectedPaths) {
      if (!dir.isDir) {
        Encrypt(dir.path);
      }
    }
  };

  const handlePathSelection = (dir: dir.Dir, isChecked: boolean) => {
    if (isChecked) {
      setSelectedPaths([...selectedPaths, dir]);
    } else {
      setSelectedPaths(selectedPaths.filter((d) => d.path !== dir.path));
    }
  };

  return (
    <PathContext.Provider
      value={{
        dirList,
        setDirList,
        paths,
        setPaths,
        selectedPaths,
        setSelectedPaths,
        handlePath,
        getDirs,
      }}
    >
      <SideBar>
        <CheckboxGroup>
          {selectedPaths.length > 0 && (
            <Button onClick={handleSelected} colorScheme="blue">
              Encrpyt selected
            </Button>
          )}
          <Grid mt={70} templateColumns="repeat(8, 1fr)">
            {dirList?.map((dir) => (
              <GridItem key={dir.path}>
                <VStack>
                  <Box
                    onDoubleClick={() =>
                      dir.isDir
                        ? handlePath(dir.path, dir.name)
                        : OpenFile(dir.path)
                    }
                    maxW={100}
                    wordBreak={"break-word"}
                  >
                    <HStack>
                      <Checkbox
                        onChange={(event) =>
                          handlePathSelection(dir, event.target.checked)
                        }
                      ></Checkbox>
                      <VStack>
                        <DirIcon dir={dir} />
                        <Text>{dir.name}</Text>
                      </VStack>
                    </HStack>
                  </Box>
                </VStack>
              </GridItem>
            ))}
          </Grid>
        </CheckboxGroup>
      </SideBar>
    </PathContext.Provider>
  );
};

export default App;

const DirIcon = ({ dir }: { dir: dir.Dir }) => {
  return dir.isDir ? <FcFolder size={60} /> : <FcFile size={60} />;
};
