// ForJWT Auth
/*import { getUserFromJwtAuth } from "./helper/AuthHelper";
import {
  useJWTAuth,
  useJWTAuthActions,
} from "../services/auth/jwt-auth/JWTAuthProvider";

export const useAuthUser = () => {
  const { user, isAuthenticated, isLoading } = useJWTAuth();
  return {
    isLoading,
    isAuthenticated,
    user: getUserFromJwtAuth(user),
  };
};

export const useAuthMethod = () => {
  const { signInUser, signUpUser, logout } = useJWTAuthActions();

  return {
    signInUser,
    logout,
    signUpUser,
  };
};*/
//For Firebase Auth
import {
  useFirebase,
  useFirebaseActions,
} from "../services/auth/firebase/FirebaseAuthProvider";
import { getUserFromFirebase } from "./helper/AuthHelper";

export const useAuthUser = () => {
  const { user, isAuthenticated, isLoading } = useFirebase();
  return {
    isLoading,
    isAuthenticated,
    user: getUserFromFirebase(user),
  };
};

export const useAuthMethod = () => {
  const {
    signInWithEmailAndPassword,
    createUserWithEmailAndPassword,
    signInWithPopup,
    logout,
  } = useFirebaseActions();

  return {
    signInWithEmailAndPassword,
    createUserWithEmailAndPassword,
    signInWithPopup,
    logout,
  };
};

//For AWS Auth
// import {getUserFromAWS} from './helper/AuthHelper';
// import {
//   useAwsCognito,
//   useAwsCognitoActions,
// } from '../services/auth/aws-cognito/AWSAuthProvider';
//
// export const useAuthUser = () => {
//   const {user, isAuthenticated, isLoading} = useAwsCognito();
//   return {
//     isLoading,
//     isAuthenticated,
//     user: getUserFromAWS(user),
//   };
// };
//
// export const useAuthMethod = () => {
//   const {
//     signIn,
//     signUpCognitoUser,
//     confirmCognitoUserSignup,
//     logout,
//   } = useAwsCognitoActions();
//
//   return {
//     signIn,
//     signUpCognitoUser,
//     confirmCognitoUserSignup,
//     logout,
//   };
// };

/*
//For Auth0
export const useAuthUser = () => {
  const {user, isAuthenticated, isLoading} = useAuth0();
  console.log(
    'user, isAuthenticated, isLoading',
    user,
    isAuthenticated,
    isLoading,
  );
  return {
    isLoading,
    isAuthenticated,
    user: useMemo(() => getUserFromAuth0(user), []),
  };
};

export const useAuthMethod = () => {
  const {loginWithRedirect, logout} = useAuth0();

  return loginWithRedirect;
};*/
