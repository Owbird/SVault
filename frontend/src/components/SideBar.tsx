import {
  Box,
  BoxProps,
  Breadcrumb,
  BreadcrumbItem,
  BreadcrumbLink,
  CloseButton,
  Flex,
  FlexProps,
  HStack,
  Icon,
  IconButton,
  Link,
  Text,
  useColorModeValue,
  useDisclosure,
} from "@chakra-ui/react";
import { ReactNode, ReactText, useContext } from "react";
import { IconType } from "react-icons";
import { BiChevronRight } from "react-icons/bi";
import { CiVault } from "react-icons/ci";
import { FiArrowLeft, FiHome, FiMenu } from "react-icons/fi";
import { PathContext } from "../contexts/pathsContext";

interface LinkItemProps {
  name: string;
  icon: IconType;
}
const LinkItems: Array<LinkItemProps> = [
  { name: "Home", icon: FiHome },
  { name: "Vault", icon: CiVault },
];

export default function SideBar({ children }: { children: ReactNode }) {
  const { isOpen, onOpen, onClose } = useDisclosure();

  const pathData = useContext(PathContext);

  const { paths, getDirs, setPaths } = pathData || {};
  return (
    <Box minH="100vh" bg={useColorModeValue("gray.100", "gray.900")}>
      <SidebarContent
        onClose={() => onClose}
        display={{ base: "none", md: "block" }}
      />

      {/* mobilenav */}
      <MobileNav
        setPaths={setPaths}
        getDirs={getDirs}
        paths={paths}
        onOpen={onOpen}
      />
      <Box ml={{ base: 0, md: 60 }} p="4">
        {children}
      </Box>
    </Box>
  );
}

interface SidebarProps extends BoxProps {
  onClose: () => void;
}

const SidebarContent = ({ onClose }: SidebarProps) => {
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
        <CloseButton display={{ base: "flex", md: "none" }} onClick={onClose} />
      </Flex>
      {LinkItems.map((link) => (
        <NavItem key={link.name} icon={link.icon}>
          {link.name}
        </NavItem>
      ))}
    </Box>
  );
};

interface NavItemProps extends FlexProps {
  icon: IconType;
  children: ReactText;
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

interface MobileProps extends FlexProps {
  onOpen: () => void;
  paths: string[];
  setPaths: (paths: string[]) => void;
  getDirs: (path: string) => void;
}
const MobileNav = ({ onOpen, paths, getDirs, setPaths }: MobileProps) => {
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
    <Flex
      pos={"fixed"}
      zIndex={1}
      w={"full"}
      ml={{ base: 0, md: 60 }}
      px={{ base: 4, md: 4 }}
      height="20"
      alignItems="center"
      bg={useColorModeValue("white", "gray.900")}
      borderBottomWidth="1px"
      borderBottomColor={useColorModeValue("gray.200", "gray.700")}
      justifyContent={{ base: "space-between", md: "flex-start" }}
    >
      <IconButton
        display={{ base: "flex", md: "none" }}
        onClick={onOpen}
        variant="outline"
        aria-label="open menu"
        icon={<FiMenu />}
      />

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
    </Flex>
  );
};
