import React from 'react';
import HomeIcon from '@material-ui/icons/Home';
import AccountCircleSharpIcon from '@material-ui/icons/AccountCircleSharp';
//import AssignmentSharpIcon from '@material-ui/icons/AssignmentSharp';
import SignOut from '@material-ui/icons/Settings';
//import AssignmentIndIcon from '@material-ui/icons/AssignmentInd';


import {
  Sidebar,
  SidebarItem,
  SidebarDivider,
  SidebarSpace,
  //SidebarUserSettings,
  //SidebarThemeToggle,
  SidebarPinButton,
} from '@backstage/core';

export const AppSidebar = () => (
  <Sidebar>
    <SidebarDivider />
    {/* Global nav, not org-specific */}
    <SidebarItem icon={HomeIcon} to="welcome" text="Home" />
    {/* <SidebarItem icon={CreateComponentIcon} to="create" text="Create..." />
    <SidebarItem icon={CreateComponentIcon} to="welcome" text="Welcome" /> */}
    
    <SidebarItem
      icon={AccountCircleSharpIcon}
      to="course"
      text="Create Course"
    />
    
    <SidebarItem
      icon={AccountCircleSharpIcon}
      to="courseclass"
      text="Create Courseclass"
    />
    
    <SidebarItem
      icon={AccountCircleSharpIcon}
      to="subjectsoffered"
      text="Create Subjects Offered"
    />

    {/* End global nav */}
    <SidebarDivider />
    <SidebarSpace />
    <SidebarDivider />
    <SidebarItem
      icon={SignOut}
      to="sign_out"
      text="Sign Out"
    />
    {/* <SidebarUserSettings  /> */}
    <SidebarPinButton />
  </Sidebar>
);