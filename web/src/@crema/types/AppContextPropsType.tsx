import {
  FooterType,
  LayoutType,
  NavStyle,
  RouteTransition,
  ThemeMode,
  ThemeStyle,
} from "../shared/constants/AppEnums";
import { LanguageProps } from "../core/AppLngSwitcher/data";
import { Breakpoints, Direction, Theme, Transitions } from "@mui/material";
import { ComponentsProps } from "@mui/material/styles/props";
import { Shadows } from "@mui/material/styles/shadows";
import { Palette } from "@mui/material/styles/createPalette";
import { Mixins } from "@mui/material/styles/createMixins";
import { ZIndex } from "@mui/material/styles/zIndex";

interface CremaPalette extends Palette {
  background: {
    paper: string;
    default: string;
  };
  primary: {
    light: string;
    main: string;
    dark: string;
    contrastText: string;
  };
  secondary: {
    light: string;
    main: string;
    dark: string;
    contrastText: string;
  };
  sidebar: {
    bgColor: string;
    textColor: string;
  };
  text: {
    primary: string;
    secondary: string;
    disabled: string;
    hint: string;
  };
  common: {
    white: string;
    black: string;
  };
  gray: {
    50: string;
    100: string;
    200: string;
    300: string;
    400: string;
    500: string;
    600: string;
    700: string;
    800: string;
    900: string;
    A100: string;
    A200: string;
    A400: string;
    A700: string;
  };
}

export interface CremaTheme extends Theme {
  direction: Direction;
  cardRadius: number;
  palette: CremaPalette;
  status: {
    danger: string;
  };
  divider: string;
  overrides: {
    MuiTypography: {
      h1: {
        fontSize: number;
      };
      h2: {
        fontSize: number;
      };
      h3: {
        fontSize: number;
      };
      h4: {
        fontSize: number;
      };
      h5: {
        fontSize: number;
      };
      h6: {
        fontSize: number;
      };
      subtitle1: {
        fontSize: number;
      };
      subtitle2: {
        fontSize: number;
      };
      body1: {
        fontSize: number;
      };
      body2: {
        fontSize: number;
      };
    };
    MuiToggleButton: {
      root: {
        borderRadius: number;
      };
    };
    MuiCardLg: {
      root: {
        borderRadius: number;
      };
    };
    MuiCard: {
      root: {
        borderRadius: number;
      };
    };
    MuiButton: {
      root: {
        borderRadius: number;
      };
    };
  };
  breakpoints: Breakpoints;
  mixins: Mixins;
  props?: ComponentsProps;
  shadows: Shadows;
  transitions: Transitions;
  zIndex: ZIndex;
  unstable_strictMode?: boolean;
}

export default interface AppContextPropsType {
  theme: CremaTheme;
  routes: any;
  themeStyle: ThemeStyle;
  themeMode: ThemeMode;
  navStyle: NavStyle;
  layoutType: LayoutType;
  footerType: FooterType;
  rtAnim: RouteTransition;
  footer: boolean;
  locale: {
    languageId: string;
    locale: string;
    name: string;
    icon: string;
  };
  rtlLocale: string[];
  primary?: string;
  secondary?: string;
  isRTL?: boolean;
  sidebarColor?: string;
  // routes,
  updateLayoutStyle?: (layoutType: LayoutType) => void;
  setRTL: (rtl: boolean) => void;
  updateSidebarColor?: (sidebarColor: string) => void;
  setFooter?: (footer: boolean) => void;
  setFooterType?: (footerType: FooterType) => void;
  updateThemeStyle?: (themeStyle: ThemeStyle) => void;
  updateTheme?: (theme: any) => void;
  updateMode?: (themeMode: ThemeMode) => void;
  updateThemeMode: (themeMode: ThemeMode) => void;
  updatePrimaryColor?: (primaryColor: string) => void;
  updateSecondaryColor?: (secondaryColor: string) => void;
  changeLocale: (locale: LanguageProps) => void;
  changeNavStyle: (navStyle: NavStyle) => void;
  changeRTAnim?: (routeTransition: RouteTransition) => void;
}
