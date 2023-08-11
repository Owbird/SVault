import { Box, useColorModeValue } from "@chakra-ui/react";
import { ReactNode } from "react";
import { IconType } from "react-icons";
import { CiVault } from "react-icons/ci";
import { FiHome } from "react-icons/fi";
import NavItem from "./NavItem";
import PathBar from "./PathBar";
import SidebarContent from "./SidebarContent";

interface LinkItemProps {
  name: string;
  icon: IconType;
}
const links: Array<LinkItemProps> = [
  { name: "Home", icon: FiHome },
  { name: "Vault", icon: CiVault },
];

export default function SideBar({ children }: { children: ReactNode }) {
  const linkItems = links.map((link) => (
    <NavItem key={link.name} name={link.name} icon={link.icon}>
      {link.name}
    </NavItem>
  ));

  return (
    <Box minH="100vh" bg={useColorModeValue("gray.100", "gray.900")}>
      <SidebarContent>{linkItems}</SidebarContent>
      <PathBar />
      <Box ml={{ base: 0, md: 60 }} p="4">
        {children}
      </Box>
    </Box>
  );
}
