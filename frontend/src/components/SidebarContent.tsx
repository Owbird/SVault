import { Box, Flex, Text, useColorModeValue } from "@chakra-ui/react";
import { FC, ReactNode } from "react";

interface SidebarContentProps {
  children: ReactNode;
}

const SidebarContent: FC<SidebarContentProps> = ({ children }) => {
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
      </Flex>
      {children}
    </Box>
  );
};

export default SidebarContent;
