import { Route } from '@angular/router';

import { isLoggedInGuard } from '@cubeshares/services/auth';

import { LayoutComponent } from '@cubeshares/components/layout/layout.component';

import { MePageComponent } from './me/me.page';
import { UserDetailsPageComponent } from './user-details/user-details.page';

export const USER_ROUTES: Route[] = [
  {
    path: '',
    component: LayoutComponent,
    children: [
      {
        path: '',
        pathMatch: 'full',
        redirectTo: 'me',
      },
      {
        path: 'me',
        canActivate: [isLoggedInGuard],
        component: MePageComponent,
        data: { requiresAuth: true },
      },
      {
        path: ':id',
        component: UserDetailsPageComponent,
      },
    ],
  },
];
