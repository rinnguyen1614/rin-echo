import React from "react";
import ContentLoader from "react-content-loader";

export const MailDetailSkeleton = () => (
  <ContentLoader
    speed={2}
    backgroundColor="#f3f3f3"
    height={600}
    foregroundColor="#ecebeb"
  >
    <rect x="100" y="35" rx="3" ry="3" width="188" height="12" />
    <rect x="100" y="55" rx="3" ry="3" width="152" height="10" />
    <rect x="105" y="106" rx="3" ry="3" width="510" height="8" />
    <rect x="105" y="132" rx="3" ry="3" width="480" height="8" />
    <rect x="105" y="158" rx="3" ry="3" width="510" height="8" />
    <rect x="105" y="178" rx="3" ry="3" width="478" height="8" />
    <rect x="105" y="198" rx="3" ry="3" width="510" height="8" />
    <rect x="105" y="218" rx="3" ry="3" width="478" height="8" />
    <circle cx="50" cy="50" r="25" />
  </ContentLoader>
);
