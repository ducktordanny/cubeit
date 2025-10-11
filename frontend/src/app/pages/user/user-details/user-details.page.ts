import { Component, effect, untracked } from '@angular/core';
import { Router } from '@angular/router';

import { injectParams } from 'ngxtension/inject-params';

import { UsersService } from '@cubeshares/services/user/users.service';
import { UserMeService } from '@cubeshares/services/user';

import { UserDetailsCardComponent } from '../components/user-details-card/user-details-card.component';
import { UserPostListComponent } from "../components/user-post-list/user-post-list.component";

@Component({
  selector: 'cubeshares-user-details-page',
  templateUrl: 'user-details.page.html',
  styleUrl: '../user.page.scss',
  imports: [UserDetailsCardComponent, UserPostListComponent],
  providers: [UsersService]
})
export class UserDetailsPageComponent {
  protected readonly user = this.usersService.user;

  private readonly id = injectParams('id');

  constructor(private readonly usersService: UsersService, private readonly userMeService: UserMeService, private readonly router: Router) {
    this.observeIdParamChanges();
  }

  private observeIdParamChanges(): void {
    effect(() => {
      const id = this.id();
      const isLoading = this.userMeService.isLoading()
      untracked(() => {
        if (id === null || isLoading) return;
        if (+id === this.userMeService.loggedInUser()?.id) {
          void this.router.navigate(['/user/me'])
          return;
        }
        this.usersService.pollReadUserById(+id);
      })
    })
  }
}
