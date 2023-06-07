import {
  Box,
  Button,
  CheckboxGroup,
  Grid,
  GridItem,
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

const App = () => {
  const [dirList, setDirList] = useState<dir.Dir[]>();
  const [paths, setPaths] = useState<string[]>([]);
  const [selectedPaths, setSelectedPaths] = useState<dir.Dir[]>([]);

  const handleSelected = () => {
    for (let dir of selectedPaths) {
      if (!dir.isDir) {
        Encrypt(dir.path);
      }
    }
  };

  const getDirs = (path: string) => {
    GetDirs(path).then(setDirList);
    setSelectedPaths([]);
  };

  const handlePath = (path: string, dir: string) => {
    getDirs(path);
    setPaths([...paths, dir]);
  };

  const handlePathSelection = (dir: dir.Dir, isChecked: boolean) => {
    if (isChecked) {
      setSelectedPaths([...selectedPaths, dir]);
    } else {
      setSelectedPaths(selectedPaths.filter((d) => d.path !== dir.path));
    }
  };

  useEffect(() => {
    GetUserHome().then((path) => handlePath(path, path));
  }, []);

  return (
    <SideBar getDirs={getDirs} setPaths={setPaths} paths={paths}>
      {selectedPaths.length > 0 && (
        <Button onClick={handleSelected} colorScheme="blue">
          Encrpyt selected
        </Button>
      )}

      <CheckboxGroup>
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
                  {/* <Checkbox
                    onChange={(event) =>
                      handlePathSelection(dir, event.target.checked)
                    }
                  > */}
                  <DirIcon dir={dir} />
                  <Text>{dir.name}</Text>
                  {/* </Checkbox> */}
                </Box>
              </VStack>
            </GridItem>
          ))}
        </Grid>
      </CheckboxGroup>
    </SideBar>
  );
};

export default App;

const DirIcon = ({ dir }: { dir: dir.Dir }) => {
  return dir.isDir ? <FcFolder size={60} /> : <FcFile size={60} />;
};
