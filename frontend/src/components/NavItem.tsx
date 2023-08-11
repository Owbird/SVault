import { Flex, FlexProps, Icon, Link } from "@chakra-ui/react";
import { ReactNode, useContext } from "react";
import { IconType } from "react-icons";
import { GetUserHome } from "../../wailsjs/go/uifunctions/UIFunctions";
import { PathContext } from "../contexts/pathsContext";

interface NavItemProps extends FlexProps {
  icon: IconType;
  name: string;
  children: ReactNode;
}

const NavItem = ({ icon, children, name }: NavItemProps) => {
  const pathContext = useContext(PathContext);

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
    <Link
      href="#"
      onClick={() => toggleBodyView(name)}
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

export default NavItem;
