import React from 'react';
import HomeIcon from '@material-ui/icons/Home';
import TableChartIcon from '@material-ui/icons/TableChart';
import SignOut from '@material-ui/icons/Settings';
import BallotIcon from '@material-ui/icons/Ballot';
import SubjectIcon from '@material-ui/icons/Subject';
//import AccountCircleSharpIcon from '@material-ui/icons/AccountCircleSharp';

import {
  Sidebar,
  SidebarItem,
  SidebarDivider,
  SidebarSpace,
  //SidebarUserSettings,
  //SidebarThemeToggle,
  SidebarPinButton,
} from '@backstage/core';

export const AppSidebar = ({setSession}) => {

  function LinklogOut(){
    setSession ({
      isLoggedIn : false,
      isSignIn : false
    });
  }

  return (
  <Sidebar>
    <SidebarDivider />
    {/* Global nav, not org-specific */}
    <SidebarItem icon={HomeIcon} to="welcome" text="Home" />
    {/* <SidebarItem icon={CreateComponentIcon} to="create" text="Create..." />
    <SidebarItem icon={CreateComponentIcon} to="welcome" text="Welcome" /> */}
    
    <SidebarItem
      icon={BallotIcon}
      to="course"
      text="Create Course"
    />
    
    <SidebarItem
      icon={TableChartIcon}
      to="courseclass"
      text="Create Class"
    />
    
    <SidebarItem
      icon={SubjectIcon}
      to="subjectsoffered"
      text="Create Subjects Offered"
    />

    {/* End global nav */}
    <SidebarDivider />
    <SidebarSpace />
    <SidebarDivider />
    <SidebarItem
      icon={SignOut}
      onClick = {LinklogOut}
      to="Login"
      text="Log Out"
    />
    {/* <SidebarUserSettings  /> */}
    <SidebarPinButton />
  </Sidebar>
  )};