import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import Login from '../../../packages/app/src/components/Login';
import Course from './components/CourseBuild';
import Courseclass from './components/Courseclass';
import EntSubjectsOffered from './components/SubjectsOffered'
//import Sign from './components/Sign-in';

export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', WelcomePage);
    router.registerRoute('/welcome', WelcomePage);
    router.registerRoute('/course', Course);
    router.registerRoute('/courseclass', Courseclass);
    router.registerRoute('/subjectsoffered', EntSubjectsOffered);
    router.registerRoute('/Login', Login);
    //router.registerRoute('/Sign-in', Sign);
  },
});
