import {
  Box,
  Breadcrumb,
  BreadcrumbItem,
  BreadcrumbLink,
  Button,
  CheckboxGroup,
  Grid,
  GridItem,
  HStack,
  Text,
  VStack,
} from "@chakra-ui/react";
import { Fragment, useEffect, useState } from "react";
import { DefaultExtensionType } from "react-file-icon";
import { BiChevronRight, BiLeftArrow } from "react-icons/bi";
import { dir } from "../wailsjs/go/models";
import {
  Encrypt,
  GetDirs,
  GetUserHome,
  OpenFile,
} from "../wailsjs/go/uifunctions/UIFunctions";

import { FcFile, FcFolder } from "react-icons/fc";

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
        {paths.length > 1 && <BiLeftArrow onClick={goBack} />}

        <Breadcrumb separator={<BiChevronRight size={30} color="gray.500" />}>
          {paths.map((path) => {
            const isActive = path === paths[paths.length - 1];
            return (
              <BreadcrumbItem
                color={`${isActive && "white"}`}
                bgColor={`${isActive && "green"}`}
                padding={`${isActive && "5px"}`}
              >
                <BreadcrumbLink onClick={() => handlePathClick(path)}>
                  {path}
                </BreadcrumbLink>
              </BreadcrumbItem>
            );
          })}
        </Breadcrumb>
      </HStack>
      {selectedPaths.length > 0 && (
        <Button onClick={handleSelected} colorScheme="blue">
          Encrpyt selected
        </Button>
      )}

      <CheckboxGroup>
        <Grid templateColumns="repeat(8, 1fr)">
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
    </Fragment>
  );
};

export default App;

const DirIcon = ({ dir }: { dir: dir.Dir }) => {
  const pathData = dir.path.split(".");

  const ext = pathData[pathData.length - 1] as DefaultExtensionType;

  if (dir.isDir) return <FcFolder size={60} />;
  return <FcFile size={60} />;
};
