import { Badge, Grid, GridItem, HStack, Text, VStack } from "@chakra-ui/react";
import { Fragment, useEffect, useState } from "react";
import { FaArrowLeft } from "react-icons/fa";
import { FcFile, FcFolder } from "react-icons/fc";
import { GetDirs, GetUserHome, OpenFile } from "../wailsjs/go/main/App";
import { main } from "../wailsjs/go/models";

const App = () => {
  const [dirList, setDirList] = useState<main.DirList[]>();
  const [paths, setPaths] = useState<string[]>([]);

  const handlePathClick = (path: string) => {
    if (paths.length > 1 && paths[paths.length - 1] !== path) {
      const newPaths = paths.slice(0, paths.indexOf(path) + 1);
      setPaths(newPaths);
      GetDirs(newPaths.join("/")).then(setDirList);
    }
  };

  const goBack = () => {
    paths.pop();

    setPaths(paths);
    GetDirs(paths.join("/")).then(setDirList);
  };

  const handlePath = (path: string, dir: string) => {
    GetDirs(path).then(setDirList);
    setPaths([...paths, dir]);
  };

  useEffect(() => {
    GetUserHome().then((path) => handlePath(path, path));
  }, []);

  return (
    <Fragment>
      <HStack>
        {paths.length > 1 && <FaArrowLeft onClick={goBack} />}
        {paths.map((path) => (
          <Badge
            onClick={() => handlePathClick(path)}
            padding={5}
            background="green"
          >
            {path}
          </Badge>
        ))}
      </HStack>
      <Grid templateColumns="repeat(7, 1fr)" gap={2}>
        {dirList?.map((dir) => (
          <GridItem
            onClick={() =>
              dir.isDir ? handlePath(dir.path, dir.name) : OpenFile(dir.path)
            }
            key={dir.path}
          >
            <VStack>
              {dir.isDir ? <FcFolder size={60} /> : <FcFile size={60} />}
              <Text>{dir.name}</Text>
            </VStack>
          </GridItem>
        ))}
      </Grid>
    </Fragment>
  );
};

export default App;
