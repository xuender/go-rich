import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { Routes, RouterModule } from '@angular/router';

import { IonicModule } from '@ionic/angular';

import { TagsPage } from './tags.page';
import { TagPageModule } from './tag/tag.module';

const routes: Routes = [
  {
    path: '',
    component: TagsPage
  }
];

@NgModule({
  imports: [
    CommonModule,
    FormsModule,
    IonicModule,
    RouterModule.forChild(routes),
    TagPageModule,
  ],
  declarations: [TagsPage]
})
export class TagsPageModule {}
