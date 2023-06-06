import {
  Badge,
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
import { Fragment, useEffect, useState } from "react";
import { FaArrowLeft } from "react-icons/fa";
import { FcFile, FcFolder } from "react-icons/fc";
import { dir } from "../wailsjs/go/models";
import {
  Encrypt,
  GetDirs,
  GetUserHome,
  OpenFile,
} from "../wailsjs/go/uifunctions/UIFunctions";

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

  const handlePathClick = (path: string) => {
    if (paths.length > 1 && paths[paths.length - 1] !== path) {
      const newPaths = paths.slice(0, paths.indexOf(path) + 1);
      setPaths(newPaths);
      getDirs(newPaths.join("/"));
    }
  };

  const goBack = () => {
    paths.pop();

    setPaths(paths);
    getDirs(paths.join("/"));
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
    <Fragment>
      <HStack>
        {paths.length > 1 && <FaArrowLeft onClick={goBack} />}
        {paths.map((path) => (
          <Badge onClick={() => handlePathClick(path)} background="green">
            {path}
          </Badge>
        ))}
      </HStack>
      {selectedPaths.length > 0 && (
        <Button onClick={handleSelected} colorScheme="blue">
          Encrpyt selected
        </Button>
      )}

      <CheckboxGroup>
        <Grid templateColumns="repeat(7, 1fr)" gap={2}>
          {dirList?.map((dir) => (
            <GridItem key={dir.path}>
              <VStack>
                <Box
                  onDoubleClick={() =>
                    dir.isDir
                      ? handlePath(dir.path, dir.name)
                      : OpenFile(dir.path)
                  }
                >
                  <Checkbox
                    onChange={(event) =>
                      handlePathSelection(dir, event.target.checked)
                    }
                  >
                    {dir.isDir ? <FcFolder size={60} /> : <FcFile size={60} />}
                    <Text>{dir.name}</Text>
                  </Checkbox>
                </Box>
              </VStack>
            </GridItem>
          ))}
        </Grid>
      </CheckboxGroup>
    </Fragment>
  );
};

export default App;
