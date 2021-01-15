import { createPlugin } from '@backstage/core';
import Sign from './Sign-in';
import Login from './Login';

export const plugin = createPlugin({
  id: 'home',
  register({ router }) {
    router.registerRoute('/Sign-in', Sign);
    router.registerRoute('/Login', Login);
  },
});