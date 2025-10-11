import { Component } from '@angular/core';

import { ButtonModule } from 'primeng/button';

import { UserMeService } from '@cubeshares/services/user';

import { UserDetailsCardComponent } from '../components/user-details-card/user-details-card.component';
import { PostCreateComponent } from "../components/post-create/post-create.component";
import { UserPostListComponent } from "../components/user-post-list/user-post-list.component";

@Component({
  selector: 'cubeshares-me-page',
  templateUrl: 'me.page.html',
  styleUrl: '../user.page.scss',
  imports: [ButtonModule, UserDetailsCardComponent, PostCreateComponent, UserPostListComponent],
})
export class MePageComponent {
  protected readonly user = this.userMeService.loggedInUser;

  constructor(private readonly userMeService: UserMeService) { }
}
