import React from 'react';
import { render } from '@testing-library/react';
import WelcomePage from './WelcomePage';
import { ThemeProvider } from '@material-ui/core';
import { darkTheme } from '@backstage/theme';

describe('WelcomePage', () => {
  it('should render', () => {
    const rendered = render(
      <ThemeProvider theme={darkTheme}>
        <WelcomePage />
      </ThemeProvider>,
    );
    expect(rendered.baseElement).toBeInTheDocument();
  });
});
