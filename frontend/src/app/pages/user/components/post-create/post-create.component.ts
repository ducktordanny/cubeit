import { Component } from "@angular/core";

import { ButtonModule } from "primeng/button";
import { CardModule } from 'primeng/card';
import { TextareaModule } from 'primeng/textarea';

@Component({
  selector: 'cubeshares-post-create',
  templateUrl: 'post-create.component.html',
  styleUrl: 'post-create.component.scss',
  imports: [ButtonModule, CardModule, TextareaModule],
})
export class PostCreateComponent { }
