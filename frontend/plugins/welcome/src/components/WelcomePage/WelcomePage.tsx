import React, { FC } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import {
  Typography,
  Grid,
  List,
  ListItem,
  ListItemText,
  Link,
} from '@material-ui/core';
import Timer from '../Timer';
import {
  Content,
  InfoCard,
  Header,
  Page,
  pageTheme,
  ContentHeader,
  SupportButton,
} from '@backstage/core';

const WelcomePage: FC<{}> = () => {
  const profile = { givenName: '' };

  return (
    <Page theme={pageTheme.home}>
      <Header
        title={`Welcome ${profile.givenName || 'to Register System'}`}
        subtitle="Let quick create what you want now."
      >
        <Timer />
      </Header>
      <Content>
        <ContentHeader title="Getting Started">
          <SupportButton />
        </ContentHeader>
        <Grid container>
          <Grid item xs={12} md={6}>
            <InfoCard>
              <Typography variant="body1" gutterBottom>
                <span role="img" aria-label="confetti">
                </span>
              </Typography>
              <Typography variant="h6" gutterBottom>
                The First!
              </Typography>
              <Typography variant="body1" paragraph>
                You can creat Course.
              </Typography>
              <Typography variant="h6" gutterBottom>
                Second ^o^
              </Typography>
              <Typography variant="body1" paragraph>
                You can creat Class.
              </Typography>
              <Typography variant="h6" gutterBottom>
                Third +.+
              </Typography>
              <Typography variant="body1" paragraph>
                You can creat Subjects Offered.
                <Link href="https://github.com/spotify/backstage/blob/master/docs/getting-started/create-a-plugin.md">
                </Link>{' '}
                {' '}
                <Link component={RouterLink} to="/home">
                  
                </Link>{' '}
                
              </Typography>
            </InfoCard>
          </Grid>
          <Grid item>
            <InfoCard>
              <Typography variant="h5">Quick Links</Typography>
              <List>
                <ListItem>
                  <Link href="https://backstage.io">backstage.io</Link>
                </ListItem>
                <ListItem>
                  <Link href="https://github.com/spotify/backstage/blob/master/docs/getting-started/create-a-plugin.md">
                    Create a plugin
                  </Link>
                </ListItem>
              </List>
            </InfoCard>
          </Grid>
        </Grid>
      </Content>
    </Page>
  );
};

export default WelcomePage;
