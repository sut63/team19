import React, { FC } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import {makeStyles} from '@material-ui/core/styles';
import Warning from "@material-ui/icons/Warning";
import TableChartIcon from '@material-ui/icons/TableChart';
import BallotIcon from '@material-ui/icons/Ballot';
import SubjectIcon from '@material-ui/icons/Subject';
import FindInPageIcon from '@material-ui/icons/FindInPage';
import LibraryBooksIcon from '@material-ui/icons/LibraryBooks';
import NotesIcon from '@material-ui/icons/Notes';
import TocIcon from '@material-ui/icons/Toc';
import HelpIcon from '@material-ui/icons/Help';
import AccountCircleIcon from '@material-ui/icons/AccountCircle';
import GridContainer from "plugin-welcome/src/support/Grid/GridContainer.js";
import GridItem from "plugin-welcome/src/support/Grid/GridItem";
import Card from "plugin-welcome/src/support/Card/Card";
import CardHeader from "plugin-welcome/src/support/Card/CardHeader";
import CardIcon from "plugin-welcome/src/support/Card/CardIcon";
import Danger from "plugin-welcome/src/support/Typography/Danger.js";
import Info from "plugin-welcome/src/support/Typography/Info";
import CardFooter from "plugin-welcome/src/support/Card/CardFooter.js";
import CardBody from "plugin-welcome/src/support/Card/CardBody";
import Table from "plugin-welcome/src/support/Table/Table";
import styleses from "plugin-welcome/src/assets/jss/material-dashboard-react/views/dashboardStyle";
import {
  Button,
  Grid,
} from '@material-ui/core';
import {
  Content,
  Page,
  pageTheme,
  ContentHeader,
} from '@backstage/core';

const useStyles = makeStyles(styleses);

const WelcomePage: FC<{}> = () => {
  const classes = useStyles();
  const blank = " ";
  const name = JSON.parse(String(localStorage.getItem("Name")));
  const title = JSON.parse(String(localStorage.getItem("Title")));

  return (
    <Page theme={pageTheme.home}>
      {/*<div className={classes.rootgid}>
      <Grid container spacing={1}>
        <Grid item xs={'auto'}>
        <div style={{ marginLeft: 7, color: 'white',}}><AccountCircleIcon/></div>
        </Grid>
        <Grid item xs={'auto'}>
        <div style={{ marginLeft: 8, color: 'white' }}>{title}{" "}{name}</div>
        </Grid>
      </Grid>
      </div>*/}
      <Content>
      <CardHeader color="info" stats icon>
      <CardIcon color="info">
        <ContentHeader title = "Welcome : Register System">
            <div style={{ marginLeft: 1, width: '800px'}}></div>
            <div style={{ marginRight: 1, color: 'white' }}><AccountCircleIcon/></div>
            <div style={{ marginRight: 1, color: 'white' }}>{title}{" "}{name}</div>
        </ContentHeader>
        </CardIcon>
        </CardHeader>
        <Grid container>
        <GridItem xs={12} sm={6} md={4}>
          <Card> 
            <CardHeader color="rose" stats icon>
              <CardIcon color="rose">
                <BallotIcon/>     
              </CardIcon>
              <p className={classes.cardCategory}>Click to build your Course now!</p>
              <h3 className={classes.cardTitle}>
              <Button
                 style={{outline: 'none', border: 'none', background: 'none'}}
                 component={RouterLink}
                 to="/Course"
               >
                 Create Course
               </Button>
              </h3>
            </CardHeader>
            <CardFooter stats>
              <div className={classes.stats}>
                <Info>
                  <HelpIcon />
                </Info> 
                  Also You can search course 👍 !!!   
              </div>
            </CardFooter>
          </Card>
          </GridItem>
          <GridItem xs={12} sm={6} md={4}>
          <Card>
            <CardHeader color="primary" stats icon>
              <CardIcon color="primary">
                <TableChartIcon/>
              </CardIcon>
              <p className={classes.cardCategory}>Click to build your Class now!</p>
              <h3 className={classes.cardTitle}>
              <Button
                 style={{outline: 'none', border: 'none', background: 'none'}}
                 component={RouterLink}
                 to="/Courseclass"
               >
                Create Class
              </Button>
              </h3> 
            </CardHeader>
            <CardFooter stats>
              <div className={classes.stats}>
              <Info>
                  <HelpIcon />
                </Info>
                Also You can search class 👌 !!!
              </div>
            </CardFooter>
          </Card>
        </GridItem>
        <GridItem xs={12} sm={6} md={4}>
          <Card>
            <CardHeader color="danger" stats icon>
              <CardIcon color="danger">
                <SubjectIcon/>
              </CardIcon>
              <p className={classes.cardCategory}>Click to build your Subjects Offered now!</p>
              <h3 className={classes.cardTitle}>
              <Button
                 style={{outline: 'none', border: 'none', background: 'none'}}
                 component={RouterLink}
                 to="/subjectsoffered"
               >
                Create Subjects Offered
              </Button>
              </h3>
            </CardHeader>
            <CardFooter stats>
              <div className={classes.stats}>
              <Info>
                  <HelpIcon />
                </Info>
                Yeah! Also You can search class ... 🖖
              </div>
            </CardFooter>
          </Card>
        </GridItem>
 
        <GridItem xs={12} sm={6} md={3}>
          <Card>
            <CardHeader color="info" stats icon>
              <CardIcon color="info">
                <FindInPageIcon/>
              </CardIcon>
              <p className={classes.cardCategory}>Click to search Instructor</p>
              <h3 className={classes.cardTitle}>
              <Button
                 style={{outline: 'none', border: 'none', background: 'none'}}
                 component={RouterLink}
                 to="/SearchInstructor"
               >
                Search Instructor
              </Button>
              </h3>
            </CardHeader>
            <CardFooter stats>
              <div className={classes.stats}>
              <Info>
                  <HelpIcon />
                </Info>
                  If you haven't apply Account, You can make it. 🤗
              </div>
            </CardFooter>
          </Card>
          </GridItem>
          <GridItem xs={12} sm={6} md={3}>
          <Card>
            <CardHeader color="success" stats icon>
              <CardIcon color="success">
                <LibraryBooksIcon/>
              </CardIcon>
              <p className={classes.cardCategory}>Click to search Course</p>
              <h3 className={classes.cardTitle}>
                <Button
                 style={{outline: 'none', border: 'none', background: 'none'}}
                 component={RouterLink}
                 to="/FindCourse"
               >
                Search Course
              </Button>
              </h3>
            </CardHeader>
            <CardFooter stats>
              <div className={classes.stats}>
              <Info>
                  <HelpIcon />
                </Info>
                Do you have Problem to find Course? Find it ! 😆
              </div>
            </CardFooter>
          </Card>
        </GridItem>
        <GridItem xs={12} sm={6} md={3}>
          <Card>
            <CardHeader color="rose" stats icon>
              <CardIcon color="rose">
                <NotesIcon/>
              </CardIcon>
              <p className={classes.cardCategory}>Click to search Subject</p>
              <h3 className={classes.cardTitle}>
              <Button
                 style={{outline: 'none', border: 'none', background: 'none'}}
                 component={RouterLink}
                 to="/searchsubjectsoffered"
               >
                Search Subject
              </Button>
              </h3>
            </CardHeader>
            <CardFooter stats>
              <div className={classes.stats}>
              <Info>
                  <HelpIcon />
                </Info>
                Oh! You lost your Subject? Find it here. 🧐
              </div>
            </CardFooter>
          </Card>
        </GridItem>
        <GridItem xs={12} sm={6} md={3}>
          <Card>
            <CardHeader color="warning" stats icon>
              <CardIcon color="warning">
                <TocIcon/>
              </CardIcon>
              <p className={classes.cardCategory}>Click to search Class</p>
              <h3 className={classes.cardTitle}>
              <Button
                 style={{outline: 'none', border: 'none', background: 'none'}}
                 component={RouterLink}
                 to=""
               >
                Search Class
              </Button>
              </h3>
            </CardHeader>
            <CardFooter stats>
              <div className={classes.stats}>
              <Info>
                  <HelpIcon />
                </Info>
                Waiting For Updated 😞 
              </div>
            </CardFooter>
          </Card>
        </GridItem>
        <GridItem xs={12} sm={12} md={12}>
          <Card>
            <CardHeader color="primary">
              <h4 className={classes.cardTitleWhite}>Survival Member Now 😅 </h4>
              <p className={classes.cardCategoryWhite}>
                Update on 5th February, 2021 
              </p>
            </CardHeader>
            <CardBody>
              <Table
                tableHeaderColor="warning"
                tableHead={["ID", "Name", "Responsibility", "Github"]}
                tableData={[
                  ["B6103637", "KRIT KAEWSATHITPORNCHAI","Instructor", "MasVorx"],
                  ["B6107192", "PATTAPHONG SAPMAK","Subject", "B6107192"],
                  ["B6108519", "PHANUWAT NUCHANRAM","Class", "B6108519"],
                  ["B6115708", "SAIKWAN SOOKKHUM","Course", "kwan3010"]
                ]}
              />
            </CardBody>
          </Card>
        </GridItem>
        </Grid>
      </Content>
    </Page>
  );
};

export default WelcomePage;
