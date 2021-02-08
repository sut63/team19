import React from 'react';
import HomeIcon from '@material-ui/icons/Home';
import TableChartIcon from '@material-ui/icons/TableChart';
import SignOut from '@material-ui/icons/Settings';
import BallotIcon from '@material-ui/icons/Ballot';
import SubjectIcon from '@material-ui/icons/Subject';
import FindInPageIcon from '@material-ui/icons/FindInPage';
import LibraryBooksIcon from '@material-ui/icons/LibraryBooks';
import NotesIcon from '@material-ui/icons/Notes';
import TocIcon from '@material-ui/icons/Toc';
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
    localStorage.setItem("ID", JSON.stringify(null));
    localStorage.setItem("Title", JSON.stringify(null));
    localStorage.setItem("Name", JSON.stringify(null));
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

    <SidebarItem
      icon={FindInPageIcon}
      to="SearchInstructor"
      text="Search Instructor"
    />

    <SidebarItem
      icon={LibraryBooksIcon}
      to="FindCourse"
      text="Search Course"
    />

    <SidebarItem
      icon={NotesIcon}
      to="searchsubjectsoffered"
      text="Search Subjects Offered"
    />

    <SidebarItem
      icon={TocIcon}
      //to=""
      text="Search Class"
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