import { Component, NgZone } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { ModalController, ActionSheetController, List } from '@ionic/angular';
import { Observable } from 'rxjs';
import { URL } from '../../api/init'
import { Tag } from '../../api/tag';
import { TagPage } from './tag/tag.page';
import { pull } from 'lodash'

@Component({
  selector: 'app-tags',
  templateUrl: './tags.page.html',
  styleUrls: ['./tags.page.scss'],
})
export class TagsPage {

  tags$: Observable<Tag[]>
  tags: Tag[] = []
  constructor(
    private http: HttpClient,
    private actionSheetCtrl: ActionSheetController,
    private ngZone: NgZone,
    private modalCtrl: ModalController
  ) {
    this.load()
  }

  load() {
    this.tags$ = this.http.get<Tag[]>(`${URL}/api/tags`)
  }

  cancel() {
    this.modalCtrl.dismiss();
  }

  has(tags: Tag[]): boolean {
    if (tags) {
      this.tags = tags
      return tags.length > 0
    }
    return false;
  }

  async create() {
    const modal = await this.modalCtrl.create({
      component: TagPage,
      componentProps: {
        tag: { name: ''}
      }
    })
    modal.onDidDismiss().then(d => {
      if (d.data) {
        this.tags.push(d.data)
      }
    })
    return await modal.present()
  }

  async update(t: Tag) {
    const modal = await this.modalCtrl.create({
      component: TagPage,
      componentProps: { tag: Object.assign({}, t) },
    });
    modal.onDidDismiss().then(d => {
      if (d.data) {
        Object.assign(t, d.data)
      }
    })
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
          il.closeSlidingItems()
          this.http.delete(`${URL}/api/tags/${t.id}`)
            .subscribe((r) => {
              pull(this.tags, t)
              this.ngZone.run(() => false)
            })
        }
      }, {
        text: '取消',
        icon: 'close',
        role: 'cancel',
      }]
    });
    await actionSheet.present();
  }
  search(event: CustomEvent) {
    const txt = event.detail.value
    this.tags$ = this.http.get<Tag[]>(`${URL}/api/tags?search=${txt}`)
  }
}
