import { CommonModule } from "@angular/common";
import { NgModule, CUSTOM_ELEMENTS_SCHEMA } from "@angular/core";

import { ConcatPipe } from "./concat.pipe";
import { ToNamePipe } from './to-name.pipe'

const pipes = [
  ConcatPipe,
  ToNamePipe,
]
@NgModule({
  imports: [
    CommonModule
  ],
  declarations: [...pipes,],
  exports: [...pipes],
  schemas: [CUSTOM_ELEMENTS_SCHEMA]
})
export class PipeModule { }
