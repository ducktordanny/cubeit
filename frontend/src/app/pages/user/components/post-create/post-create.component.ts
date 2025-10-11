import { Component } from "@angular/core";
import { FormBuilder, ReactiveFormsModule } from "@angular/forms";

import { ButtonModule } from "primeng/button";
import { CardModule } from 'primeng/card';
import { TextareaModule } from 'primeng/textarea';

@Component({
  host: {
    '(keydown.meta.enter)': 'onSubmit()',
    '(keydown.control.enter)': 'onSubmit()',
  },
  selector: 'cubeshares-post-create',
  templateUrl: 'post-create.component.html',
  styleUrl: 'post-create.component.scss',
  imports: [ButtonModule, CardModule, ReactiveFormsModule, TextareaModule],
})
export class PostCreateComponent {
  protected form = this.fb.group({
    description: this.fb.control(''),
  });

  constructor(private readonly fb: FormBuilder) { }

  protected onSubmit(): void {
    console.log(this.form.getRawValue());
  }

  protected onClear(): void {
    this.form.reset();
  }
}
