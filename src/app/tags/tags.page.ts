import { Component } from '@angular/core';
import { ModalController, ActionSheetController, List } from '@ionic/angular';
import { Observable } from 'rxjs';

import { Tag } from './tag';
import { TagPage } from './tag/tag.page';
import { TagService } from './tag.service';

@Component({
  selector: 'app-tags',
  templateUrl: './tags.page.html',
  styleUrls: ['./tags.page.scss'],
})

export class TagsPage {
  tags$: Observable<Tag[]>
  title: string = ''
  constructor(
    public tagService: TagService,
    private actionSheetCtrl: ActionSheetController,
    private modalCtrl: ModalController
  ) {
    this.tags$ = this.tagService.tags$
  }

  cancel() {
    this.modalCtrl.dismiss();
  }

  async create() {
    const modal = await this.modalCtrl.create({
      component: TagPage,
      componentProps: {
        tag: { name: '' }
      }
    })
    return await modal.present()
  }

  async update(t: Tag) {
    const modal = await this.modalCtrl.create({
      component: TagPage,
      componentProps: { tag: Object.assign({}, t) },
    });
    return await modal.present()
  }

  async del(t: Tag, il: List) {
    const actionSheet = await this.actionSheetCtrl.create({
      header: `确认删除标签 [ ${t.name} ] ?`,
      buttons: [{
        text: '删除',
        role: 'destructive',
        icon: 'trash',
        handler: () => {
          this.tagService.delete(t)
        }
      }, {
        text: '取消',
        icon: 'close',
        role: 'cancel',
      }]
    });
    await actionSheet.present();
    await il.closeSlidingItems()
  }

  search(event: CustomEvent) {
    const txt = event.detail.value
    this.title = `搜索: ${txt}`
    this.tags$ = this.tagService.search(txt)
  }
}
