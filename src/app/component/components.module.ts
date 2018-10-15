import { FormsModule } from '@angular/forms';
import { IonicModule, } from '@ionic/angular';
import { CommonModule } from '@angular/common';
import { NgModule, CUSTOM_ELEMENTS_SCHEMA } from '@angular/core';

import { ExtComponent } from '../ext/ext.component';
import { TagsComponent } from '../tags/tags.component';
import { NameComponent } from './name/name.component';

const components = [
  TagsComponent,
  ExtComponent,
  NameComponent,
]
@NgModule({
  declarations: [...components],
  imports: [
    CommonModule,
    FormsModule,
    IonicModule,
  ],
  exports: [...components],
  schemas: [CUSTOM_ELEMENTS_SCHEMA],
})
export class ComponentsModule { }
