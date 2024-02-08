import { extendTheme } from "@chakra-ui/react";
import { mode } from "@chakra-ui/theme-tools";

export const Theme = extendTheme({
  styles: {
    global: (props: any) => ({
      body: {
        bg: mode("blue.50", "grey")(props)
      },
    }),
  },
});