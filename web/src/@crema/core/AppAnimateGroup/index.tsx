// import React, { memo, useEffect } from "react";
// import Grow, { GrowProps } from "@mui/material/Grow";
// import Fade, { FadeProps } from "@mui/material/Fade";
// import Slide, { SlideProps } from "@mui/material/Slide";
// import Zoom, { ZoomProps } from "@mui/material/Zoom";
// import Collapse, { CollapseProps } from "@mui/material/Collapse";
//
// export type AnimType = "grow" | "fade" | "slide" | "zoom" | "collapse";
// type AnimProps = GrowProps | FadeProps | SlideProps | ZoomProps | CollapseProps;
// type RequireAtLeastOne<T, Keys extends keyof T = keyof T> = Pick<
//   T,
//   Exclude<keyof T, Keys>
// > &
//   {
//     [K in Keys]-?: Required<Pick<T, K>> & Partial<Pick<T, Exclude<Keys, K>>>;
//   }[Keys];
//
// function usePrevious(value: any) {
//   const ref = React.useRef();
//   useEffect(() => {
//     ref.current = value;
//   });
//   return ref.current || [];
// }
//
// interface ItemProps {
//   shown: boolean;
//   children: any;
//   onCompleteOutAnimation?: VoidFunction;
//   onExited: VoidFunction;
//   timeout?: { enter?: number; exit?: number };
//   animation: AnimType;
//   animationProps?: AnimProps;
// }
//
// function AnimatedItem({
//   shown,
//   children,
//   timeout,
//   onExited,
//   animationProps,
//   animation,
// }: ItemProps) {
//   useEffect(() => {}, [shown]);
//   const componentMap: any = {
//     grow: Grow,
//     fade: Fade,
//     slide: Slide,
//     zoom: Zoom,
//     collapse: Collapse,
//   };
//   const SelectedComponent = componentMap[animation];
//   return (
//     <SelectedComponent
//       {...animationProps}
//       timeout={timeout}
//       in={shown}
//       onExiting={onExited}>
//       {children}
//     </SelectedComponent>
//   );
// }
//
// interface AppAnimateGroupProps {
//   children: RequireAtLeastOne<any, "key">[] | RequireAtLeastOne<any, "key">;
//   animation?: AnimType;
//   animationProps?: AnimProps;
//   initialAnimationDuration?: number;
// }
//
// const AppAnimateGroup = ({
//   children,
//   animation = "grow",
//   animationProps,
//   initialAnimationDuration = 750,
// }: AppAnimateGroupProps) => {
//   const previousChildren: any = usePrevious(children);
//   const [removed, setRemoved] = React.useState<{ [index: number]: any }>([]);
//   const [removedShown, setRemovedShown] = React.useState<{
//     [index: number]: any;
//   }>([]);
//
//   useEffect(() => {
//     const removeChildren = () => {
//       const newlyRemoved = previousChildren.filter(
//         (c: any) => children.findIndex((oc: any) => oc.key === c.key) === -1
//       );
//       newlyRemoved.forEach((r: any) => {
//         const index = previousChildren.findIndex((rr: any) => r.key === rr.key);
//         setRemoved({ ...removed, [index]: r });
//         setRemovedShown({ ...removedShown, [index]: r });
//         setTimeout(() => {
//           delete removedShown[index];
//           setRemovedShown({ ...removedShown });
//         }, 100);
//       });
//     };
//
//     if (previousChildren.length > children.length) {
//       removeChildren();
//     }
//   }, [children, previousChildren, removedShown, removed]);
//
//   const handleExit = (index: any) => {
//     setTimeout(() => {
//       delete removed[index];
//       setRemoved({ ...removed });
//     }, 300);
//   };
//
//   const getEnterDelayTime = (index: number) => {
//     return initialAnimationDuration * ((index + 1) / (children.length || 1));
//   };
//   return (
//     <>
//       {children.length === 0 && removed[0] ? (
//         <AnimatedItem
//           onExited={() => handleExit(0)}
//           key={removed[0].key}
//           shown={removedShown[0] !== undefined}
//           timeout={{ enter: 0 }}
//           animation={animation}
//           animationProps={animationProps}>
//           {removed[0]}
//         </AnimatedItem>
//       ) : (
//         children.map((Child: any, i: number) => (
//           <React.Fragment key={i}>
//             {i === 0 && removed[i] && (
//               <AnimatedItem
//                 animation={animation}
//                 onExited={() => handleExit(i)}
//                 key={removed[i].key}
//                 shown={removedShown[i] !== undefined}
//                 timeout={{ enter: 0, exit: 200 }}>
//                 {removed[i]}
//               </AnimatedItem>
//             )}
//             <AnimatedItem
//               animation={animation}
//               animationProps={animationProps}
//               shown={true}
//               key={Child.key || i}
//               onExited={() => handleExit(Child.key)}
//               timeout={{
//                 enter: previousChildren.find((p: any) => p.key === Child.key)
//                   ? 0
//                   : getEnterDelayTime(i),
//               }}>
//               {Child}
//             </AnimatedItem>
//             {removed[i + 1] && (
//               <AnimatedItem
//                 animation={animation}
//                 animationProps={animationProps}
//                 onExited={() => handleExit(i + 1)}
//                 key={removed[i + 1].key}
//                 shown={removedShown[i + 1] !== undefined}
//                 timeout={{ enter: 0, exit: 500 }}>
//                 {removed[i + 1]}
//               </AnimatedItem>
//             )}
//           </React.Fragment>
//         ))
//       )}
//     </>
//   );
// };
//
// export default memo(AppAnimateGroup);

import React, { ReactNode } from "react";

interface AppAnimateGroupProps {
  children: ReactNode;

  [x: string]: any;
}

const AppAnimateGroup: React.FC<AppAnimateGroupProps> = (props) => {
  return <div style={props.style}>{props.children}</div>;
};

export default AppAnimateGroup;
