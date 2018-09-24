import { NgModule, CUSTOM_ELEMENTS_SCHEMA } from '@angular/core';
import { IonicModule, } from '@ionic/angular';
import { TagsComponent } from '../tags/tags.component';
import { CommonModule } from '@angular/common';
import { ExtComponent } from '../ext/ext.component';
import { FormsModule } from '@angular/forms';

const components = [
  TagsComponent,
  ExtComponent,
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
