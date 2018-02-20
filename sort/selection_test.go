package goalgorithms

import (
	"reflect"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	tests := []struct {
		name string
		list []int
		want []int
	}{
		{"Mixed", []int{1, 3, 4, 5, 2, 9, 8, 0}, []int{0, 1, 2, 3, 4, 5, 8, 9}},
		{"Already sorted", []int{0, 1, 2, 3, 4, 5, 8, 9}, []int{0, 1, 2, 3, 4, 5, 8, 9}},
		{"Almost sorted", []int{0, 1, 2, 3, 4, 5, 9, 8}, []int{0, 1, 2, 3, 4, 5, 8, 9}},
		{"Reversed", []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SelectionSort(tt.list)
			got := tt.list
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectionSort(%v) = %v, want %v", tt.list, got, tt.want)
			}
		})
	}
}

func BenchmarkSelectionSort(b *testing.B) {
	a := []int{266, 393, 134, 490, 464, 643, 463, 654, 177, 424, 347, 389, 859, 17, 455, 640, 771, 819, 577, 865, 976, 59, 532, 857, 868, 607, 27, 79, 560, 945, 833, 12, 123, 592, 908, 133, 938, 390, 156, 259, 319, 912, 686, 636, 318, 545, 950, 204, 632, 246, 469, 384, 8, 367, 629, 731, 16, 146, 655, 194, 610, 544, 441, 523, 961, 432, 479, 190, 376, 212, 700, 189, 719, 42, 161, 781, 218, 214, 468, 536, 371, 793, 825, 817, 73, 92, 575, 303, 422, 94, 769, 321, 435, 173, 26, 129, 900, 809, 762, 206, 642, 767, 127, 937, 729, 552, 844, 835, 226, 831, 310, 964, 712, 6, 97, 512, 802, 621, 641, 7, 339, 141, 498, 669, 453, 822, 618, 858, 798, 704, 362, 445, 517, 944, 568, 74, 920, 187, 529, 959, 270, 525, 45, 89, 947, 984, 290, 208, 199, 628, 921, 500, 923, 186, 958, 906, 138, 61, 365, 242, 19, 980, 328, 527, 239, 681, 936, 209, 348, 329, 181, 783, 896, 154, 722, 132, 446, 240, 410, 847, 645, 737, 680, 293, 599, 874, 811, 766, 377, 258, 118, 866, 854, 387, 427, 755, 286, 625, 349, 496, 878, 824, 595, 505, 466, 14, 230, 70, 953, 675, 41, 967, 423, 495, 483, 638, 316, 450, 31, 883, 457, 535, 996, 236, 184, 501, 888, 222, 727, 442, 276, 1000, 197, 796, 113, 924, 903, 882, 356, 713, 342, 667, 55, 305, 593, 160, 46, 195, 459, 977, 541, 178, 344, 898, 553, 904, 381, 516, 589, 829, 379, 658, 485, 317, 804, 36, 144, 602, 795, 412, 673, 157, 689, 373, 773, 927, 570, 957, 690, 263, 300, 845, 11, 679, 968, 23, 790, 249, 499, 201, 108, 816, 77, 437, 803, 966, 741, 970, 633, 941, 57, 37, 702, 696, 448, 846, 994, 416, 894, 343, 867, 95, 351, 447, 90, 652, 96, 13, 528, 340, 547, 677, 164, 213, 202, 614, 555, 691, 782, 5, 877, 826, 784, 635, 497, 843, 674, 683, 56, 873, 323, 578, 473, 569, 998, 579, 355, 631, 398, 706, 150, 744, 418, 910, 777, 392, 738, 67, 697, 736, 283, 322, 962, 298, 302, 391, 739, 986, 733, 982, 606, 223, 665, 312, 624, 336, 971, 53, 506, 278, 969, 345, 368, 522, 460, 279, 334, 761, 991, 477, 255, 488, 508, 183, 576, 383, 313, 661, 480, 746, 491, 881, 233, 692, 400, 648, 147, 850, 911, 346, 725, 420, 444, 889, 324, 413, 718, 808, 82, 511, 35, 561, 776, 827, 143, 69, 465, 671, 139, 153, 252, 772, 155, 335, 538, 353, 550, 10, 472, 588, 372, 990, 787, 973, 694, 605, 405, 751, 651, 519, 159, 556, 337, 102, 720, 284, 591, 15, 540, 760, 946, 685, 884, 895, 103, 601, 234, 84, 436, 609, 98, 879, 358, 955, 193, 327, 666, 603, 38, 730, 93, 40, 49, 88, 949, 682, 370, 905, 928, 997, 66, 956, 264, 241, 269, 594, 734, 531, 880, 99, 801, 415, 574, 291, 600, 216, 251, 431, 925, 341, 254, 582, 111, 307, 85, 1, 458, 482, 639, 34, 267, 364, 647, 438, 664, 818, 244, 748, 979, 922, 378, 830, 622, 231, 308, 325, 128, 789, 814, 554, 530, 703, 221, 745, 315, 122, 68, 119, 705, 314, 176, 117, 275, 124, 109, 112, 486, 247, 676, 487, 855, 292, 542, 559, 564, 875, 581, 714, 893, 224, 238, 24, 909, 852, 940, 114, 104, 548, 245, 235, 918, 33, 101, 740, 22, 919, 121, 105, 382, 406, 926, 662, 823, 871, 207, 261, 402, 732, 566, 375, 86, 620, 933, 812, 174, 585, 395, 539, 39, 617, 110, 282, 47, 974, 562, 567, 604, 543, 834, 891, 584, 65, 586, 452, 484, 885, 649, 71, 672, 295, 942, 149, 797, 916, 426, 476, 434, 18, 297, 573, 659, 260, 220, 62, 580, 756, 815, 272, 929, 670, 185, 778, 369, 169, 289, 481, 897, 248, 820, 791, 507, 513, 972, 217, 571, 288, 137, 331, 612, 870, 203, 598, 716, 167, 219, 215, 408, 747, 503, 975, 630, 848, 851, 876, 887, 172, 256, 188, 764, 309, 428, 489, 4, 792, 504, 821, 572, 768, 757, 750, 148, 841, 902, 421, 754, 140, 860, 419, 993, 853, 253, 44, 999, 533, 985, 232, 63, 698, 107, 708, 326, 752, 78, 76, 701, 965, 724, 366, 352, 839, 794, 397, 474, 440, 182, 627, 663, 191, 800, 872, 165, 549, 100, 478, 237, 180, 211, 992, 83, 136, 115, 917, 439, 842, 939, 759, 862, 403, 320, 262, 277, 3, 205, 626, 563, 456, 838, 200, 934, 449, 914, 60, 425, 493, 467, 9, 509, 799, 650, 989, 131, 91, 770, 721, 684, 142, 735, 583, 805, 608, 715, 863, 723, 228, 864, 168, 374, 280, 510, 717, 471, 417, 828, 557, 443, 120, 832, 492, 758, 166, 558, 81, 409, 886, 411, 171, 952, 470, 125, 951, 514, 350, 742, 126, 837, 892, 668, 462, 285, 130, 546, 257, 657, 728, 265, 960, 357, 524, 587, 145, 597, 304, 710, 116, 404, 354, 306, 330, 152, 869, 225, 20, 360, 29, 932, 765, 749, 786, 399, 454, 596, 359, 243, 25, 840, 634, 943, 948, 981, 653, 52, 333, 502, 294, 613, 611, 274, 250, 401, 915, 930, 565, 849, 615, 890, 763, 551, 268, 43, 414, 656, 363, 807, 388, 162, 726, 135, 775, 361, 433, 810, 271, 54, 526, 699, 534, 281, 332, 451, 151, 21, 616, 58, 590, 301, 711, 175, 385, 806, 394, 644, 637, 198, 196, 80, 338, 273, 106, 520, 983, 813, 396, 28, 210, 87, 678, 386, 30, 978, 2, 785, 688, 179, 32, 287, 429, 774, 537, 913, 687, 963, 407, 709, 660, 780, 743, 163, 693, 779, 461, 515, 64, 931, 299, 51, 788, 227, 707, 170, 430, 380, 518, 836, 646, 494, 50, 192, 988, 861, 901, 229, 72, 695, 48, 995, 296, 623, 753, 935, 75, 158, 475, 856, 619, 987, 907, 954, 521, 311, 899}
	for i := 0; i < b.N; i++ {
		SelectionSort(a)
	}
}
