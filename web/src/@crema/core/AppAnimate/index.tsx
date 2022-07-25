import React, { memo } from "react";
import { VelocityComponent, VelocityComponentProps } from "velocity-react";
import "velocity-animate/velocity.ui";

interface AppAnimateProps extends VelocityComponentProps {
  visibility?: string;
  duration?: number;
  delay?: number;
  easing?: any;
  display?: any;
  children: any;

  [x: string]: any;
}

const AppAnimate: React.FC<AppAnimateProps> = ({
  animation = "transition.fadeIn",
  runOnMount = true,
  targetQuerySelector = null,
  interruptBehavior = "stop",
  visibility = "visible",
  duration = 400,
  delay = 100,
  easing = [0.4, 0.0, 0.2, 1],
  display = null,
  children,
  rest,
}) => {
  const child = React.cloneElement(children, {
    style: {
      ...children.style,
      visibility: "hidden",
    },
  });
  return (
    <VelocityComponent
      animation={animation}
      runOnMount={runOnMount}
      targetQuerySelector={targetQuerySelector}
      interruptBehavior={interruptBehavior}
      visibility={visibility}
      duration={duration}
      delay={delay}
      easing={easing}
      display={display}
      {...rest}
    >
      {child}
    </VelocityComponent>
  );
};

export default memo(AppAnimate);
