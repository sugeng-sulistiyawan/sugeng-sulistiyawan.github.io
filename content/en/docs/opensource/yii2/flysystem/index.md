---
title: "Yii2 Flysystem"
description: "The League Flysystem for local and remote filesystems library for Yii2."
lead: "The League Flysystem for local and remote filesystems library for Yii2."
date: 2023-05-05T02:08:30+07:00
lastmod: 2023-05-05T02:08:33+07:00
draft: false
images: []
menu:
  docs:
    parent: "opensource-yii2"
weight: 20
toc: true
---

> [https://github.com/sugeng-sulistiyawan/yii2-flysystem](https://github.com/sugeng-sulistiyawan/yii2-flysystem)

---

This extension provides [Flysystem 3](https://flysystem.thephpleague.com) integration for the Yii framework.
[Flysystem](https://flysystem.thephpleague.com) is a filesystem abstraction which allows you to easily swap out a local filesystem for a remote one.

> Yii2 Flysystem uses [league/flysystem](https://github.com/thephpleague/flysystem)

## Instalation

Package is available on [Packagist](https://packagist.org/packages/diecoding/yii2-flysystem), you can install it using [Composer](https://getcomposer.org).

```shell
composer require diecoding/yii2-flysystem "^1.0"
```

or add to the require section of your `composer.json` file.

```shell
"diecoding/yii2-flysystem": "^1.0"
```

## Dependencies

- PHP 8.0+
- [yiisoft/yii2](https://github.com/yiisoft/yii2)
- [league/flysystem](https://github.com/thephpleague/flysystem)

## Dev. Dependencies

- [league/flysystem-aws-s3-v3](https://github.com/thephpleague/flysystem-aws-s3-v3)
- [league/flysystem-ftp](https://github.com/thephpleague/flysystem-ftp)

## Configuration

### Local Filesystem

Configure application `components` as follows

```php
return [
    // ...
    'components' => [
        // ...
        'fs' => [
            'class'  => \diecoding\flysystem\LocalComponent::class,
            'path'   => dirname(dirname(__DIR__)) . '/storage', // or you can use @alias
            'key'    => 'my-key',
            'secret' => 'my-secret',
            'action' => '/site/file', // action for get url file
            'prefix' => '',
            // 'cipherAlgo' => 'aes-128-cbc',
        ],
    ],
];
```

### AWS S3 Filesystem

Either run

```shell
composer require league/flysystem-aws-s3-v3:^3.0
```

or add

```shell
"league/flysystem-aws-s3-V3": "^3.0"
```

to the `require` section of your `composer.json` file and configure application `components` as follows

```php
return [
    // ...
    'components' => [
        // ...
        'fs' => [
            'class'    => \diecoding\flysystem\AwsS3Component::class,
            'endpoint' => 'http://your-endpoint'
            'key'      => 'your-key',
            'secret'   => 'your-secret',
            'bucket'   => 'your-bucket',
            'prefix'   => '',
            // 'region'               => 'us-east-1'
            // 'version'              => 'latest',
            // 'usePathStyleEndpoint' => false,
            // 'streamReads'          => false,
            // 'options'              => [],
            // 'credentials'          => [],
        ],
    ],
];
```

### FTP Filesystem

Either run

```shell
composer require league/flysystem-ftp:^3.0
```

or add

```shell
"league/flysystem-aws-s3-V3": "^3.0"
```

to the `require` section of your `composer.json` file and configure application `components` as follows

```php
return [
    // ...
    'components' => [
        // ...
        'fs' => [
            'class'    => \diecoding\flysystem\FTPComponent::class,
            'host'     => 'hostname',
            'root'     => '/root/path/', // or you can use @alias
            'username' => 'username',
            'password' => 'password',
            // 'port'                            => 21,
            // 'ssl'                             => false,
            // 'timeout'                         => 90,
            // 'utf8'                            => false,
            // 'passive'                         => true,
            // 'transferMode'                    => FTP_BINARY,
            // 'systemType'                      => null,         // 'windows' or 'unix'
            // 'ignorePassiveAddress'            => null,         // true or false
            // 'timestampsOnUnixListingsEnabled' => false,
            // 'recurseManually'                 => true,
            // 'useRawListOptions'               => null,         // true or false
            'action' => '/site/file', // action for get url file
            'prefix' => '',
        ],
    ],
];
```

## Additional Configuration

### URL File Action Settings

The following adapters have URL File Action generation capabilities:

- Local Component
- FTP Component

Configure `action` in `controller` as follows

> This example at `SiteController` for `/site/file`

```php
class SiteController extends Controller
{
    //...
    public function actions()
    {
        return [
            // ...
            'file' => [
                'class' => \diecoding\flysystem\actions\FileAction::class,
                // 'component' => 'fs',
            ],
        ];
    }
}
```

> Remember to configure `action` key in `fs` application components as follows

```php
return [
    // ...
    'components' => [
        // ...
        'fs' => [
            // ...
            'action' => '/site/file', // action for get url file
        ],
    ],
];
```

### Global Visibility Settings

Configure `fs` application component as follows

```php
return [
    //...
    'components' => [
        //...
        'fs' => [
            //...
            'config' => [
                'visibility' => \League\Flysystem\Visibility::PRIVATE,
            ],
        ],
    ],
];
```

## Usage

### Writing or Updating Files

To write or update file

```php
Yii::$app->fs->write('filename.ext', 'contents');
```

To write or update file using stream contents

```php
$stream = fopen('/path/to/somefile.ext', 'r+');
Yii::$app->fs->writeStream('filename.ext', $stream);
```

### Reading Files

To read file

```php
$contents = Yii::$app->fs->read('filename.ext');
```

To retrieve a read-stream

```php
$stream = Yii::$app->fs->readStream('filename.ext');
$contents = stream_get_contents($stream);
fclose($stream);
```

### Checking if a File Exists

To check if a file exists

```php
$exists = Yii::$app->fs->fileExists('filename.ext');
```

### Deleting Files

To delete file

```php
Yii::$app->fs->delete('filename.ext');
```

### Getting Files Mime Type

To get file mime type

```php
$mimeType = Yii::$app->fs->mimeType('filename.ext');
```

### Getting Files Timestamp / Last Modified

To get file timestamp

```php
$timestamp = Yii::$app->fs->lastModified('filename.ext');
```

### Getting Files Size

To get file size

```php
$byte = Yii::$app->fs->fileSize('filename.ext');
```

### Creating Directories

To create directory

```php
Yii::$app->fs->createDirectory('path/to/directory');
```

Directories are also made implicitly when writing to a deeper path

```php
Yii::$app->fs->write('path/to/filename.ext');
```

### Checking if a Directory Exists

To check if a directory exists

```php
$exists = Yii::$app->fs->directoryExists('path/to/directory');
```

### Deleting Directories

To delete directory

```php
Yii::$app->fs->deleteDirectory('path/to/directory');
```

### Checking if a File or Directory Exists

To check if a file or directory exists

```php
$exists = Yii::$app->fs->has('path/to/directory/filename.ext');
```

### Managing Visibility

Visibility is the abstraction of file permissions across multiple platforms. Visibility can be either public or private.

```php
Yii::$app->fs->write('filename.ext', 'contents', [
    'visibility' => \League\Flysystem\Visibility::PRIVATE
]);
```

You can also change and check visibility of existing files

```php
if (Yii::$app->fs->visibility('filename.ext') === \League\Flysystem\Visibility::PRIVATE) {
    Yii::$app->fs->setVisibility('filename.ext', \League\Flysystem\Visibility::PUBLIC);
}
```

### Listing contents

To list contents

```php
$contents = Yii::$app->fs->listContents();

foreach ($contents as $object) {
    echo $object['basename']
        . ' is located at' . $object['path']
        . ' and is a ' . $object['type'];
}
```

By default Flysystem lists the top directory non-recursively. You can supply a directory name and recursive boolean to get more precise results

```php
$contents = Yii::$app->fs->listContents('path/to/directory', true);
```

### Copy Files or Directories

To copy contents

```php
Yii::$app->fs->copy('path/from/directory/filename.ext', 'path/to/directory/filename.ext', [
    'visibility' => \League\Flysystem\Visibility::PRIVATE
]);
```

### Move Files or Directories

To move contents

```php
Yii::$app->fs->move('path/from/directory/filename.ext', 'path/to/directory/filename.ext', [
    'visibility' => \League\Flysystem\Visibility::PRIVATE
]);
```

### Get URL Files

To get url contents

```php
Yii::$app->fs->publicUrl('path/to/directory/filename.ext');
```

### Get URL Temporary Files / Presigned URL

To get temporary url contents

```php
$expiresAt = new \DateTimeImmutable('+10 Minutes');

Yii::$app->fs->temporaryUrl('path/to/directory/filename.ext', $expiresAt);
```

The `$expiresAt` should be a valid and instance of `PHP DateTimeInterface`. Read [PHP documentation](https://www.php.net/manual/en/class.datetimeinterface.php) for details.

### Get MD5 Hash File Contents

To get MD5 hash of the file contents

```php
Yii::$app->fs->checksum('path/to/directory/filename.ext');
```

## Using Traits

### Model Trait

Attach the Trait to the `Model/ActiveRecord` with some media attribute that will be saved in Flysystem (fs):

```php
/**
 * @property string|null $file
 */
class Model extends \yii\db\ActiveRecord
{
    use \diecoding\flysystem\traits\ModelTrait;

    // ...

    public function rules()
    {
        return [
            ['image', 'string'], // Stores the filename
        ];
    }

    /**
     * @inheritdoc
     */
    protected function attributePaths()
    {
        return [
            'image' => 'images/'
        ];
    }

    // ...
}
```

Override the `attributePaths()` method to change the base path where the files will be saved on Flysystem (fs).

- You can map a different path to each file attribute of your `Model/ActiveRecord`.

#### Using Trait Methods

```php
$image = \yii\web\UploadedFile::getInstance($model, 'image');

// Save image_thumb.* to Flysystem (fs) on //my_bucket/images/ path
// The extension of the file will be determined by the submitted file type
// This allows multiple file types upload (png,jpg,gif,...)
// $model->image will hold "image_thumb.png" after this call finish with success
$model->saveUploadedFile($image, 'image', 'image_thumb');
$model->save();

// Save image_thumb.png to Flysystem (fs) on //my_bucket/images/ path
// The extension of the file will be determined by the submitted file type
// This force the extension to *.png
$model->saveUploadedFile($image, 'image', 'image_thumb.png', false);
$model->save();

// Remove the file with named saved on the image attribute
// Continuing the example, here "//my_bucket/images/my_image.png" will be deleted from Flysystem (fs)
$model->removeFile('image');
$model->save();

// Get the URL to the image on Flysystem (fs)
$model->getFileUrl('image');

// Get the presigned URL to the image on Flysystem (fs)
// The default duration is "+5 Minutes"
$model->getFilePresignedUrl('image');
```

#### Overriding Trait Methods

##### getFsComponent

The Flysystem (fs) MediaTrait depends on this component to be configured. The default configuration is to use this component on index `'fs'`, but you may use another value. For this cases, override the `getFsComponent()` method:

```php
public function getFsComponent()
{
    return Yii::$app->get('my_fs_component');
}
```

##### attributePaths

The main method to override is `attributePaths()`, which defines a path in Flysystem (fs) for each attribute of yout model. Allowing you to save each attribute in a different Flysystem (fs) folder.

Here an example:

```php
protected function attributePaths()
{
    return [
        'logo' => 'logos/',
        'badge' => 'images/badges/'
    ];
}

// or use another attribute, example: id
// ! Note: id must contain a value first if you don't want it to be empty

protected function attributePaths()
{
    return [
        'logo' => 'thumbnail/' . $this->id . '/logos/',
        'badge' => 'thumbnail/' . $this->id . '/images/badges/'
    ];
}
```

##### getPresignedUrlDuration

The default pressigned URL duration is set to "+5 Minutes", override this method and use your own expiration.
Return must instance of `DateTimeInterface`

```php
protected function getPresignedUrlDuration($attribute)
{
    return new \DateTimeImmutable('+2 Hours');
}

// or if you want to set the attribute differently

protected function getPresignedUrlDuration($attribute)
{
    switch ($attribute) {
        case 'badge':
            return new \DateTimeImmutable('+2 Hours');
            break;

        default:
            return new \DateTimeImmutable('+1 Days');
            break;
    }
}

```

The value should be a valid and instance of `PHP DateTimeInterface`. Read [PHP documentation](https://www.php.net/manual/en/class.datetimeinterface.php) for details.
