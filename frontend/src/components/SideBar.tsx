import {
  Box,
  Breadcrumb,
  BreadcrumbItem,
  BreadcrumbLink,
  Button,
  CloseButton,
  Flex,
  FlexProps,
  HStack,
  Icon,
  Link,
  Stack,
  Text,
  useColorModeValue,
} from "@chakra-ui/react";
import { ReactNode, useContext, useState } from "react";
import { IconType } from "react-icons";
import { BiChevronRight } from "react-icons/bi";
import { CiVault } from "react-icons/ci";
import { FiArrowLeft, FiHome } from "react-icons/fi";
import { dir } from "../../wailsjs/go/models";
import {
  Decrypt,
  DeleteFile,
  Encrypt,
  GetDirs,
  GetUserHome,
  MoveFromVault,
  MoveToVault,
} from "../../wailsjs/go/uifunctions/UIFunctions";
import { PathContext, PathData } from "../contexts/pathsContext";

interface LinkItemProps {
  name: string;
  icon: IconType;
}
const LinkItems: Array<LinkItemProps> = [
  { name: "Home", icon: FiHome },
  { name: "Vault", icon: CiVault },
];

export default function SideBar({ children }: { children: ReactNode }) {
  const pathContext = useContext(PathContext);

  return (
    <Box minH="100vh" bg={useColorModeValue("gray.100", "gray.900")}>
      <SidebarContent pathContext={pathContext} />

      {/* mobilenav */}
      <MobileNav />
      <Box ml={{ base: 0, md: 60 }} p="4">
        {children}
      </Box>
    </Box>
  );
}

const SidebarContent = ({ pathContext }: { pathContext: PathData }) => {
  const toggleBodyView = async (link: string) => {
    pathContext.setCurrentBody(link);
    if (link.toLocaleLowerCase() === "home") {
      pathContext.getDirs("/");
      pathContext.setPaths([await GetUserHome()]);
    } else {
      pathContext.getDirs(".vault");
      pathContext.setPaths([".vault"]);
    }
  };
  return (
    <Box
      transition="3s ease"
      bg={useColorModeValue("white", "gray.900")}
      borderRight="1px"
      borderRightColor={useColorModeValue("gray.200", "gray.700")}
      w={{ base: "full", md: 60 }}
      pos="fixed"
      h="full"
    >
      <Flex h="20" alignItems="center" mx="8" justifyContent="space-between">
        <Text fontSize="2xl" fontFamily="monospace" fontWeight="bold">
          SVault
        </Text>
        <CloseButton display={{ base: "flex", md: "none" }} />
      </Flex>
      {LinkItems.map((link) => (
        <NavItem
          onClick={() => toggleBodyView(link.name)}
          key={link.name}
          icon={link.icon}
        >
          {link.name}
        </NavItem>
      ))}
    </Box>
  );
};

interface NavItemProps extends FlexProps {
  icon: IconType;
  children: ReactNode;
}
const NavItem = ({ icon, children, ...rest }: NavItemProps) => {
  return (
    <Link
      href="#"
      style={{ textDecoration: "none" }}
      _focus={{ boxShadow: "none" }}
    >
      <Flex
        align="center"
        p="4"
        mx="4"
        borderRadius="lg"
        role="group"
        cursor="pointer"
        _hover={{
          bg: "cyan.400",
          color: "white",
        }}
        {...rest}
      >
        {icon && (
          <Icon
            mr="4"
            fontSize="16"
            _groupHover={{
              color: "white",
            }}
            as={icon}
          />
        )}
        {children}
      </Flex>
    </Link>
  );
};

const MobileNav = () => {
  const pathData = useContext(PathContext);

  const { selectedPaths, paths, getDirs, setPaths } = pathData || {};

  const [isEncrypting, setIsEncrypting] = useState(false);

  const encryptor = async (dirs: dir.Dir[], pwd: string) => {
    for (let dir of dirs) {
      if (dir.isDir) {
        await walkPath(dir.path, pwd);
      } else {
        await Encrypt(dir.path, pwd);

        const path_data = dir.path.split(".");

        const ovl_path =
          path_data.slice(0, path_data.length - 1).toString() + ".ovl";
        await MoveToVault(ovl_path);
        await DeleteFile(dir.path);
      }
    }
  };

  const decryptor = async (dirs: dir.Dir[], pwd: string) => {
    for (let dir of dirs) {
      if (dir.isDir) {
        await walkPath(dir.path, pwd);
      } else {
        const res = await Decrypt(dir.path, pwd);

        if (res === "Incorrect password") {
          return alert("Incorrect password");
        }
        await MoveFromVault(res);
        await DeleteFile(dir.path);
      }
    }
  };

  const walkPath = async (path: string, pwd: string) => {
    await GetDirs(path).then((dirs) =>
      pathData.currentBody === "Home"
        ? encryptor(dirs, pwd)
        : decryptor(dirs, pwd)
    );
  };

  const getPassword = (): string | null => {
    const password = prompt(
      "Enter a password or leave it blank to automatically generate one."
    );

    return password;
  };

  const handleSelected = async () => {
    setIsEncrypting(true);

    const password = getPassword();

    if (password) {
      if (pathData.currentBody == "Home") {
        await encryptor(selectedPaths, password);
      } else {
        await decryptor(selectedPaths, password);
      }
    }

    getDirs(paths.slice(0, paths.length).join("/"));

    setIsEncrypting(false);
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
  return (
    <Stack
      pos={"fixed"}
      zIndex={1}
      w={"full"}
      ml={{ base: 0, md: 60 }}
      px={{ base: 4, md: 4 }}
      height="20"
      alignItems="flex-start"
      bg={useColorModeValue("white", "gray.900")}
      borderBottomWidth="1px"
      borderBottomColor={useColorModeValue("gray.200", "gray.700")}
      justifyContent={{ base: "space-between", md: "flex-start" }}
    >
      <Text
        display={{ base: "flex", md: "none" }}
        fontSize="2xl"
        fontFamily="monospace"
        fontWeight="bold"
      >
        SVault
      </Text>

      <HStack spacing={{ base: "0", md: "6" }}>
        <HStack>{paths.length > 1 && <FiArrowLeft onClick={goBack} />}</HStack>

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
        <Button
          isLoading={isEncrypting}
          onClick={handleSelected}
          colorScheme="blue"
        >
          {pathData.currentBody === "Home" ? "Encrypt" : "Decrypt"} selected
        </Button>
      )}
    </Stack>
  );
};
